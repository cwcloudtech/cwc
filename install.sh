parse_commandline() {
  while test $# -gt 0
  do
    key="$1"
	case "$key" in
	  -i|--install-dir)
	    PARSED_INSTALL_DIR="$2"
		shift
	   ;;
	  -b|--bin-dir)
	    PARSED_BIN_DIR="$2"
		shift
	   ;;
	  -u|--update)
	    PARSED_UPGRADE="yes"
	  ;;
	  *)
	   die "Got an unexpected argument: $1"
	  ;;
    esac
	shift
  done
}

set_global_vars() {
  ROOT_INSTALL_DIR=${PARSED_INSTALL_DIR:-/usr/local/cwc}
  BIN_DIR=${PARSED_BIN_DIR:-/usr/local/bin}
  UPGRADE=${PARSED_UPGRADE:-no}

  EXE_NAME="cwc"
  INSTALLER_DIR="$( cd "$( dirname "$0" )" >/dev/null 2>&1 && pwd )"
  INSTALLER_DIST_DIR="$INSTALLER_DIR/dist"
  INSTALLER_EXE="$INSTALLER_DIST_DIR/$EXE_NAME"
  CWC_EXE_VERSION=$($INSTALLER_EXE --version | cut -d ' ' -f 1 | cut -d '/' -f 2)

  INSTALL_DIR="$ROOT_INSTALL_DIR/$CWC_EXE_VERSION"
  INSTALL_DIR="$INSTALL_DIR"
  INSTALL_DIST_DIR="$INSTALL_DIR/dist"
  INSTALL_BIN_DIR="$INSTALL_DIR/bin"
  INSTALL_CWC_EXE="$INSTALL_BIN_DIR/$EXE_NAME"

  CURRENT_INSTALL_DIR="$ROOT_INSTALL_DIR/v2/current"
  CURRENT_CWC_EXE="$CURRENT_INSTALL_DIR/bin/$EXE_NAME"

  BIN_CWC_EXE="$BIN_DIR/$EXE_NAME"
}

create_install_dir() {
  mkdir -p "$INSTALL_DIR" || exit 1
  {
    setup_install_dist &&
    setup_install_bin &&
    create_current_symlink
  } || {
    rm -rf "$INSTALL_DIR"
    exit 1
  }
}

check_preexisting_install() {
  if [ -L "$CURRENT_INSTALL_DIR" ] && [ "$UPGRADE" = "no" ]
  then
    die "Found preexisting CWC CLI installation: $CURRENT_INSTALL_DIR. Please rerun install script with --update flag."
  fi
  if [ -d "$INSTALL_DIR" ]
  then
    echo "Found same CWC CLI version: $INSTALL_DIR. Skipping install."
    exit 0
  fi
}

setup_install_dist() {
  cp -r "$INSTALLER_DIST_DIR" "$INSTALL_DIST_DIR"
}

setup_install_bin() {
  mkdir -p "$INSTALL_BIN_DIR"
  ln -s "../dist/$EXE_NAME" "$INSTALL_CWC_EXE"
}

create_current_symlink() {
  ln -snf "$INSTALL_DIR" "$CURRENT_INSTALL_DIR"
}

create_bin_symlinks() {
  mkdir -p "$BIN_DIR"
  ln -sf "$CURRENT_CWC_EXE" "$BIN_CWC_EXE"
  ln -sf "$CURRENT_CWC_COMPLETER_EXE" "$BIN_CWC_COMPLETER_EXE"
}

die() {
	err_msg="$1"
	echo "$err_msg" >&2
	exit 1
}

main() {
  parse_commandline "$@"
  set_global_vars
  check_preexisting_install
  create_install_dir
  create_bin_symlinks
  echo "You can now run: $BIN_CWC_EXE --version"
  exit 0
}

main "$@" || exit 1

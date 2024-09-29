#!/bin/bash

# Ensure the script is run as root or with sudo privileges
check_sudo() {
    if [ "$EUID" -ne 0 ]; then
        echo "This script must be run as root or with sudo privileges."
        echo "Please run the script as: sudo $0"
        exit 1
    fi
}

check_unzip() {
    if ! command -v unzip &>/dev/null; then
        echo "Unzip is not installed. Installing..."

        # Determine the package manager and install unzip
        if [[ -f /etc/debian_version ]]; then
            sudo apt-get update && sudo apt-get install -y unzip
        elif [[ -f /etc/redhat-release ]]; then
            sudo yum install -y unzip
        elif [[ -f /etc/arch-release ]]; then
            sudo pacman -S --noconfirm unzip
        else
            echo "Unsupported OS. Please install unzip manually."
            exit 1
        fi

        if ! command -v unzip &>/dev/null; then
            echo "Failed to install unzip. Please install it manually."
            exit 1
        fi

        echo "Unzip installed successfully."
    else
        echo "Unzip is already installed."
    fi
}

install_xray() {
    REPO_OWNER="XTLS"
    REPO_NAME="Xray-core"
    GITHUB_API_URL="https://api.github.com/repos/${REPO_OWNER}/${REPO_NAME}/releases/latest"
    INSTALL_DIR="/usr/local/bin/xray"
    TEMP_DIR="/tmp/xray"

    OS=$(uname | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)

    case "$ARCH" in
    x86_64)
        ARCH_PATTERN="linux-64"
        ;;
    i386 | i686)
        ARCH_PATTERN="linux-32"
        ;;
    aarch64)
        ARCH_PATTERN="arm64-v8a"
        ;;
    armv7l)
        ARCH_PATTERN="arm32-v7a"
        ;;
    loongarch64)
        ARCH_PATTERN="loong64"
        ;;
    mips64)
        ARCH_PATTERN="mips64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
    esac

    echo "Fetching latest release info for ${REPO_OWNER}/${REPO_NAME}..."
    RELEASE_INFO=$(curl -sL $GITHUB_API_URL)

    if [[ -z "$RELEASE_INFO" ]]; then
        echo "Failed to fetch release information."
        exit 1
    fi

    LATEST_TAG=$(echo "$RELEASE_INFO" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    ASSETS_URL=$(echo "$RELEASE_INFO" | grep '"assets_url":' | sed -E 's/.*"([^"]+)".*/\1/')

    echo "Latest release: $LATEST_TAG"
    echo "Fetching assets..."

    ASSETS_INFO=$(curl -sL "$ASSETS_URL")

    DOWNLOAD_URL=$(echo "$ASSETS_INFO" | grep "browser_download_url" | grep "$OS" | grep "$ARCH_PATTERN" | grep -v ".dgst" | sed -E 's/.*"([^"]+)".*/\1/')

    if [[ -z "$DOWNLOAD_URL" ]]; then
        echo "No matching asset found for OS: $OS and ARCH: $ARCH_PATTERN."
        exit 1
    fi

    mkdir -p "$TEMP_DIR"

    echo "Downloading from $DOWNLOAD_URL..."
    curl -Lo "$TEMP_DIR/xray.zip" "$DOWNLOAD_URL"

    if [[ $? -eq 0 ]]; then
        echo "Download successful."
    else
        echo "Failed to download the asset."
        exit 1
    fi

    # Ensure unzip is installed
    check_unzip

    echo "Unzipping the Xray archive into $TEMP_DIR..."
    unzip -o "$TEMP_DIR/xray.zip" -d "$TEMP_DIR"

    if [[ $? -eq 0 ]]; then
        echo "Unzipping successful."
    else
        echo "Failed to unzip the archive."
        exit 1
    fi

    # Ensure /usr/local/bin/xray directory exists
    sudo mkdir -p "$INSTALL_DIR"

    # Move xray, geoip.dat, and geosite.dat to /usr/local/bin/xray/
    echo "Moving files to $INSTALL_DIR..."
    sudo mv "$TEMP_DIR/xray" "$INSTALL_DIR/"
    sudo mv "$TEMP_DIR/geoip.dat" "$INSTALL_DIR/"
    sudo mv "$TEMP_DIR/geosite.dat" "$INSTALL_DIR/"

    if [[ $? -eq 0 ]]; then
        echo "Files moved successfully to $INSTALL_DIR."
    else
        echo "Failed to move files."
        exit 1
    fi

    # Clean up the temp directory
    rm -rf "$TEMP_DIR"

    echo "Installation completed successfully."
}

install_xrayping() {
    REPO_OWNER="galavpncom"
    REPO_NAME="xrayping"
    GITHUB_API_URL="https://api.github.com/repos/${REPO_OWNER}/${REPO_NAME}/releases/latest"
    INSTALL_DIR="/usr/local/bin/xrayping"
    TEMP_DIR="/tmp/xrayping"

    OS=$(uname | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)

    case "$ARCH" in
    x86_64)
        ARCH_PATTERN="linux-amd64"
        ;;
    i386 | i686)
        ARCH_PATTERN="linux-386"
        ;;
    aarch64)
        ARCH_PATTERN="arm64-v8a"
        ;;
    armv7l)
        ARCH_PATTERN="arm32-v7a"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
    esac

    echo "Fetching latest release info for ${REPO_OWNER}/${REPO_NAME}..."
    RELEASE_INFO=$(curl -sL $GITHUB_API_URL)

    if [[ -z "$RELEASE_INFO" ]]; then
        echo "Failed to fetch release information."
        exit 1
    fi

    LATEST_TAG=$(echo "$RELEASE_INFO" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    ASSETS_URL=$(echo "$RELEASE_INFO" | grep '"assets_url":' | sed -E 's/.*"([^"]+)".*/\1/')

    echo "Latest release: $LATEST_TAG"
    echo "Fetching assets..."

    ASSETS_INFO=$(curl -sL "$ASSETS_URL")

    DOWNLOAD_URL=$(echo "$ASSETS_INFO" | grep "browser_download_url" | grep "$OS" | grep "$ARCH_PATTERN" | grep -v ".dgst" | sed -E 's/.*"([^"]+)".*/\1/')

    if [[ -z "$DOWNLOAD_URL" ]]; then
        echo "No matching asset found for OS: $OS and ARCH: $ARCH_PATTERN."
        exit 1
    fi

    mkdir -p "$TEMP_DIR"

    echo "Downloading from $DOWNLOAD_URL..."
    curl -Lo "$TEMP_DIR/xrayping.zip" "$DOWNLOAD_URL"

    if [[ $? -eq 0 ]]; then
        echo "Download successful."
    else
        echo "Failed to download the asset."
        exit 1
    fi

    # Ensure unzip is installed
    check_unzip

    echo "Unzipping the XrayPing archive into $TEMP_DIR..."
    unzip -o "$TEMP_DIR/xrayping.zip" -d "$TEMP_DIR"

    if [[ $? -eq 0 ]]; then
        echo "Unzipping successful."
    else
        echo "Failed to unzip the archive."
        exit 1
    fi

    # Ensure /usr/local/bin/xrayping directory exists
    sudo mkdir -p "$INSTALL_DIR"

    # Move xrayping, geip.dat, and geosite.dat to /usr/local/bin/xrayping/
    echo "Moving files to $INSTALL_DIR..."
    sudo mv "$TEMP_DIR/xrayping" "$INSTALL_DIR/"

    if [[ $? -eq 0 ]]; then
        echo "Files moved successfully to $INSTALL_DIR."
    else
        echo "Failed to move files."
        exit 1
    fi

    # Clean up the temp directory
    rm -rf "$TEMP_DIR"

    echo "Installation of XrayPing completed successfully."
}

check_sudo

install_xrayping

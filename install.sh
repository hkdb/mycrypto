#!/bin/bash

PATH=$(pwd)
echo "Creating symlink at ~/.local/bin/mycrypto pointing to $PATH/mycrypto"
/usr/bin/ln -s $PATH/mycrypto ~/.local/bin

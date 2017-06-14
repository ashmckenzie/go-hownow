# Get the version.
version=`cat VERSION`
# Write out the package.
cat << EOF > version.go
package main

//go:generate bash ../scripts/get_version.sh
var version = "$version"
EOF

# Get the version.
version=`cat VERSION`
# Write out the package.
cat << EOF > version.go
package hownow

//go:generate bash ../scripts/get_version.sh

// Version ...
var Version = "$version"
EOF

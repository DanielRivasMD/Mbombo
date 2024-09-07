####################################################################################################

_default:
  @just --list

####################################################################################################

# print justfile
[group('just')]
@show:
  bat .justfile --language make

####################################################################################################

# edit justfile
[group('just')]
@edit:
  micro .justfile

####################################################################################################
# aliases
####################################################################################################

####################################################################################################
# import
####################################################################################################

# config
import '.just/go.conf'

####################################################################################################
# jobs
####################################################################################################

# build for OSX
[group('go')]
osx app=goapp:
  @echo "\n\033[1;33mBuilding\033[0;37m...\n=================================================="
  go build -v -o excalibur/{{app}}

####################################################################################################

# build for linux
[group('go')]
linux app=goapp:
  @echo "\n\033[1;33mBuilding\033[0;37m...\n=================================================="
  env GOOS=linux GOARCH=amd64 go build -v -o excalibur/{{app}}

####################################################################################################

# install locally
[group('go')]
install app=goapp exe=goexe:
  @echo "\n\033[1;33mInstalling\033[0;37m...\n=================================================="
  go install
  mv -v "${HOME}/go/bin/{{app}}" "${HOME}/go/bin/{{exe}}"

####################################################################################################

# watch changes
[group('go')]
watch:
  watchexec --clear --watch cmd -- 'just install'

####################################################################################################

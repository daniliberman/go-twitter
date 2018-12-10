#!/bin/bash

# usage: sh test.sh [-c] [-v] [-m] [-p <specific_package>]
# -c produces coverage report
# -v runs go tests with verbose option
# -p <specific_package> runs go tests only from the specified package
# -u updates dependencies for each package using dep ensure

# example: sh test.sh -v -c -p "godog"
# parses for options -c, -v, -m and -t
function opts() {
	while getopts "p:cvu" opt; do
		case $opt in
		    p) specific_package="$OPTARG"; ;;
			c) cover="true"; ;;
			v) verbose="-v"; ;;
			u) update="-u"; ;;
		esac
	done
}

opts $@

export GO_ENVIRONMENT=gokvsclient_test
export KEY_VALUE_STORE_TEST_END_POINT_READ=MOCK_READ_URL.com
export KEY_VALUE_STORE_TEST_END_POINT_WRITE=MOCK_WRITE_URL.com
export KEY_VALUE_STORE_TEST_CONTAINER_NAME=TEST

if [ $update ]; then
    current_dir=${PWD}

    for dir in $(find . -name "Gopkg.toml" -not -path "**/vendor/*" |sed 's#\(.*\)/.*#\1#' | sort -u | grep "$specific_package")
    do
        cd $dir
	dep ensure
        update_result="$?"
        cd $current_dir
        if [ $update_result != 0 ]; then
            break
        fi
    done
fi

if [ $cover ]; then
    current_dir=${PWD}

    mkdir coverage
    touch coverage/coverage_all.out
    for dir in $(find . -name "*_test.go" -not -path "**/vendor/*" |sed 's#\(.*\)/.*#\1#' | sort -u | grep "$specific_package")
    do
        cd $dir
        go test $verbose -coverprofile=$current_dir/coverage/coverage.out
        test_result="$?"
        cd $current_dir
        cat coverage/coverage.out >> coverage/coverage_all.out
        if [ $test_result != 0 ]; then
            break
        fi
    done
    awk '!a[$0]++' coverage/coverage_all.out > coverage/coverage_all2.out   #remove duplicate lines w "mode: set"
    go tool cover -html=coverage/coverage_all2.out
    rm -r coverage

else
    for dir in $(find . -name "*_test.go" -not -path "**/vendor/*" |sed 's#\(.*\)/.*#\1#' | sort -u | grep "$specific_package")
    do
           cd $dir
           go test $verbose
           test_result="$?"
           cd -
           if [ $test_result != 0 ]; then
                break
           fi
    done
fi

exit $((test_result))

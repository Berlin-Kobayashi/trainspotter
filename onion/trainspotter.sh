#!/bin/sh

DIR=$(dirname $0)

shut_off_led() {
   sh "$DIR"/expled-off.sh > /dev/null
}
shut_off_led
trap shut_off_led SIGHUP SIGINT SIGTERM

while read MESSAGE
do
    echo "$MESSAGE"

    STATUS=${MESSAGE%% *}
    BODY=${MESSAGE#* }
    case $STATUS in
        GO)
            expled 0xffffff > /dev/null
            ;;
        WAIT)
            expled 0x000000 > /dev/null
            ;;
        ERROR)
            shut_off_led
            ;;
    esac
done

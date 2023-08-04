#!/bin/bash

function pack_to_tgz(){
#!/bin/bash

# Check if the correct number of arguments is provided
# if [ $# -ne 5 ]; then
#     echo "Usage: $0 <connection_json_file> <metadata_json_file> <connection_new_value> <metadata_new_value> <package_new_file>"
#     exit 1
# fi

local connection_json_file=$1
local metadata_json_file=$2
local connection_new_value=$3
local metadata_new_value=$4
local package_new_file=$5
local new_folder=$6
# Check if the JSON file exists
if [ ! -f "$connection_json_file" ]; then
    echo "Error: File '$connection_json_file' not found."
    exit 1
fi

# Load the JSON file and update the value for key "a"
jq --arg connection_new_value "$connection_new_value" '.address = $connection_new_value' "$connection_json_file" > tmpfile.json

# Overwrite the original JSON file with the updated content
mv tmpfile.json connection.json

echo "Value of key 'address' updated to $connection_new_value in 'connection.json'."



tar cfz code.tar.gz connection.json

rm connection.json

# Check if the JSON file exists
if [ ! -f "$metadata_json_file" ]; then
    echo "Error: File '$metadata_json_file' not found."
    exit 1
fi

# Load the JSON file and update the value for key "a"
jq --arg metadata_new_value "$metadata_new_value" '.label = $metadata_new_value' "$metadata_json_file" > tmpfile.json

# Overwrite the original JSON file with the updated content
mv tmpfile.json metadata.json

echo "Value of key 'label' updated to $metadata_new_value in 'metadata.json'."

tar cfz $package_new_file code.tar.gz metadata.json
rm metadata.json code.tar.gz
mv $package_new_file $new_folder
}


mkdir basic$1

pack_to_tgz basic/packaging/connection.json basic/packaging/metadata.json basic$1-org1:7052 basic$1 basic$1-org1.tgz basic$1
pack_to_tgz basic/packaging/connection.json basic/packaging/metadata.json basic$1-org2:7052 basic$1 basic$1-org2.tgz basic$1
pack_to_tgz basic/packaging/connection.json basic/packaging/metadata.json basic$1-org3:7052 basic$1 basic$1-org3.tgz basic$1

sudo chown -R nobody:nogroup basic$1

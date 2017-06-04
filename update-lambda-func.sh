gopherjs build -m $1.go
rm -r candidate
mkdir candidate
mkdir -p candidate/node_modules/alexa-sdk/lib
mv $1.js candidate/index.js
pushd candidate
npm install alexa-sdk
zip $2.zip -r *
aws lambda update-function-code --function-name $2 --zip-file fileb://$2.zip
popd

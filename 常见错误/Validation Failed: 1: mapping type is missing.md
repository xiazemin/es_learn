https://stackoverflow.com/questions/35814080/validation-failed-1-mapping-type-is-missing-in-elasticsearch/40697210

POST index/types/_mappings
{
    "dynamic": "strict",
    "properties": {


需要指定type 
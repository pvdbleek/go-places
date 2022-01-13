```
curl http://localhost:8080/places \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "3","name": "Heaven's Gate","country": "China","description": "A stairway to heaven on Tianmen Mountain","latitude": 29.053743429510085,"longitude": 110.48154034958873}'
```
#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${GREEN}Starting load test...${NC}"

# Test 1: Basic request rate test
echo -e "\n${GREEN}Test 1: Basic request rate test${NC}"
bombardier -c 10 -n 1000 -m POST -H "Content-Type: application/json" -b '{
  "update_id": 1,
  "message": {
    "message_id": 1,
    "from": {
      "id": 123456789,
      "first_name": "Test",
      "username": "test_user",
      "is_bot": false
    },
    "chat": {
      "id": 123456789,
      "type": "private",
      "username": "test_user",
      "first_name": "Test"
    },
    "date": 1743064088,
    "text": "test message"
  }
}' http://localhost:8082/tg-updates

# Test 2: High concurrency test
echo -e "\n${GREEN}Test 2: High concurrency test${NC}"
bombardier -c 50 -n 5000 -m POST -H "Content-Type: application/json" -b '{
  "update_id": 2,
  "message": {
    "message_id": 2,
    "from": {
      "id": 123456789,
      "first_name": "Test",
      "username": "test_user",
      "is_bot": false
    },
    "chat": {
      "id": 123456789,
      "type": "private",
      "username": "test_user",
      "first_name": "Test"
    },
    "date": 1743064088,
    "text": "high load test"
  }
}' http://localhost:8082/tg-updates

# Test 3: Mixed update types test
echo -e "\n${GREEN}Test 3: Mixed update types test${NC}"

# Message update
bombardier -c 5 -n 500 -m POST -H "Content-Type: application/json" -b '{
  "update_id": 3,
  "message": {
    "message_id": 3,
    "from": {
      "id": 123456789,
      "first_name": "Test",
      "username": "test_user",
      "is_bot": false
    },
    "chat": {
      "id": 123456789,
      "type": "private",
      "username": "test_user",
      "first_name": "Test"
    },
    "date": 1743064088,
    "text": "test message"
  }
}' http://localhost:8082/tg-updates &

# Callback query update
bombardier -c 5 -n 500 -m POST -H "Content-Type: application/json" -b '{
  "update_id": 4,
  "callback_query": {
    "id": "123456789",
    "from": {
      "id": 123456789,
      "first_name": "Test",
      "username": "test_user",
      "is_bot": false
    },
    "message": {
      "message_id": 4,
      "chat": {
        "id": 123456789,
        "type": "private",
        "username": "test_user",
        "first_name": "Test"
      },
      "date": 1743064088,
      "text": "test message"
    },
    "chat_instance": "-123456789",
    "data": "test_callback"
  }
}' http://localhost:8082/tg-updates &

# Chat member update
bombardier -c 5 -n 500 -m POST -H "Content-Type: application/json" -b '{
  "update_id": 5,
  "chat_member": {
    "chat": {
      "id": 123456789,
      "type": "private",
      "username": "test_user",
      "first_name": "Test"
    },
    "from": {
      "id": 123456789,
      "first_name": "Test",
      "username": "test_user",
      "is_bot": false
    },
    "date": 1743064088,
    "old_chat_member": {
      "user": {
        "id": 987654321,
        "first_name": "Old",
        "username": "old_user",
        "is_bot": false
      },
      "status": "left"
    },
    "new_chat_member": {
      "user": {
        "id": 987654321,
        "first_name": "New",
        "username": "new_user",
        "is_bot": false
      },
      "status": "member"
    }
  }
}' http://localhost:8082/tg-updates &

# Wait for all background processes to complete
wait

# Check metrics endpoint
echo -e "\n${GREEN}Checking metrics endpoint...${NC}"
curl -s http://localhost:8082/metrics | grep -E "api_requests_total|api_request_duration_seconds|telegram_updates_total|module_requests_total"

echo -e "\n${GREEN}Load test completed.${NC}"
echo "You can now check the Grafana dashboard at http://localhost:3000" 
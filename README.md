# Chat (websocket + redis + load balancing)

Usage:

1. git clone
2. docker-compose up
3. go to localhost:8000 in the new tab
4. go to localhost:8000 in the new incognito tab
5. type messages in both tabs
6. observe that messages are synchronized between different replicas (docker logs)
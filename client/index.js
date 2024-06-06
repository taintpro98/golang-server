const EventSource = require('eventsource');

const eventSource = new EventSource('http://localhost:5000/sse/event-stream');

eventSource.onmessage = event => {
    console.log(event.data);
};
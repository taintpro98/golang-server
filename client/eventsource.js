const EventSource = require('eventsource');

const url1 = 'http://localhost:5000/v1/public/sse/newsfeed';
const url2 = 'http://localhost:5000/sse/event-stream';

const token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNTY2YjJhZTItNTgzNy00YjIwLWEwMzAtYjY4MjUzMDhjMjg4In0.LO9tOycEMp5cqRYnt3Gvvz13Rg25KqzKjBNMZkPGuszyZBgge8U7iEJp-Oq-_N5MS-HoOhGdvPCMMlcccyhDMT2UJmH8j5X28v_PlnXkXN4ae9JtAwPWxlGSjAOJpDn8GRtkfVGnC2BY4IWVZY6ulxqxFNMyTdbCPODzQXIra0kMIZpSAzirbb4gX92xLyEL05TjopDGUTXf4rC99i_NgE10eWgbZyIbdndhiBTYGgpkGg5nuRTzefYBIFV2TZ_1wCTqE3u9ob5hRj4gHEoFXRDr8AV7A_BXHWF_bjWWnHQiDp4eftU5P6qtcwQNEeK6OS0wlbItxR8OIe7R8Kg6Poh9RTxfVGLsGUrWwP3KRzH7-PBZe71r-DHpe2qB_4oscZ2l--VyI2WkSADE7L30DJM440DCqeCxJCqFb82C4b6V2Ner4KTNUghY9QCrfw2V5G8OPfe0r0jJEQzKM2EBVv4steVbCHtV3ldJ5j3jqyeEawNJMpS-iCIZuQBoE01J9fKmboZbpEsLvRZABFuZ16Lf9JSKl3U64CYjfZtlZQHANdbFqW2robbhLPyylinzK6yZOsOx4wuUc447lJcO5Kk4RDzWmQijd8w-LHz4yYfm1GAv1HndZVPYDx7-fID1-UTc3rdYbxSHrsWxCjCsObVudLn4Dzu_eJIWnvM66XY"

// Create the EventSource object with the desired URL
const eventSource = new EventSource(url1, {
    headers: {
        Authorization: `Bearer ${token}`
    }
});

// // Set the Authorization header with the bearer token
eventSource.onopen = function () {
    console.log('Connected');
};

// Handle incoming SSE events
eventSource.onmessage = function (event) {
    // Process the incoming event data
    const data = JSON.parse(event.data);
    console.log('Received event:', data);
};

// Handle SSE connection errors
eventSource.onerror = function (error) {
    console.error('SSE error:', error);
};
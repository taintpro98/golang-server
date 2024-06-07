const users = [
    {
        user_id: "566b2ae2-5837-4b20-a030-b6825308c288",
        token: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNTY2YjJhZTItNTgzNy00YjIwLWEwMzAtYjY4MjUzMDhjMjg4In0.LO9tOycEMp5cqRYnt3Gvvz13Rg25KqzKjBNMZkPGuszyZBgge8U7iEJp-Oq-_N5MS-HoOhGdvPCMMlcccyhDMT2UJmH8j5X28v_PlnXkXN4ae9JtAwPWxlGSjAOJpDn8GRtkfVGnC2BY4IWVZY6ulxqxFNMyTdbCPODzQXIra0kMIZpSAzirbb4gX92xLyEL05TjopDGUTXf4rC99i_NgE10eWgbZyIbdndhiBTYGgpkGg5nuRTzefYBIFV2TZ_1wCTqE3u9ob5hRj4gHEoFXRDr8AV7A_BXHWF_bjWWnHQiDp4eftU5P6qtcwQNEeK6OS0wlbItxR8OIe7R8Kg6Poh9RTxfVGLsGUrWwP3KRzH7-PBZe71r-DHpe2qB_4oscZ2l--VyI2WkSADE7L30DJM440DCqeCxJCqFb82C4b6V2Ner4KTNUghY9QCrfw2V5G8OPfe0r0jJEQzKM2EBVv4steVbCHtV3ldJ5j3jqyeEawNJMpS-iCIZuQBoE01J9fKmboZbpEsLvRZABFuZ16Lf9JSKl3U64CYjfZtlZQHANdbFqW2robbhLPyylinzK6yZOsOx4wuUc447lJcO5Kk4RDzWmQijd8w-LHz4yYfm1GAv1HndZVPYDx7-fID1-UTc3rdYbxSHrsWxCjCsObVudLn4Dzu_eJIWnvM66XY"
    },
    {
        user_id: "73122f85-50db-4f1d-a618-672f0609fe61",
        token: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNzMxMjJmODUtNTBkYi00ZjFkLWE2MTgtNjcyZjA2MDlmZTYxIn0.GsIyPynLypV33G39JKmPwskFPM4hf5EEJv3Of57jAz-cbZKDzooVdzhv8JRoZCG_40p4-N3cnY5pojdbep7gd4chRV-qj7hnTqce-HA-HDsxFXJ4fK-3Ce5bp9GOVqpbQjc4wCgiU_E6y4qLzaR9TXNP-vR9-xgmI6Kq-da5JyW9G8tf2U2FeoWeUz8hkiIv9fD0oO37APgzuM0I4CWT3BNOGU8rNOVzE8xyWdZ3E4uaU2d-dyQVpTi6Ci2fTpeZgck2U8TkD_3u9KjmW5SunoHWBtesrKLr6bG7TLOxXaF_ziRHE96uicxy5Y_0mJUZQPWe5jDzYDBXxri6d9vejf2u6pnfzSLv5ti9sQHj2vG1XdSJ-0akQyV1baKODrs8GgBvj8Q9udoOwjNLwVofDuoNA1RdM_Z2KMyh5AwjJTiwZiSZHMSYw_cJQYaVHf1KZjtc11CwbxoR0QrcRYvpkEexhjR78A5VGibvRxKlCs9EAu5jilkCnttSjVQ4uxwTNE4aQw79oQm4_Mke2Ucus1Vrwe4m1UueAhj8NImok4IwJ2kLyDFOzToKNLYbRpFPDvWIirSqXW5cL9je4hW35tdneI_Kb0PidaQ9cFEt7PeOTZtQnYu2WW5wtrlDYbghGcIpF4DgRXT-WrVrh3U3q1jwIZ4hEh2yQDHmx95xiSQ"
    }
]

const args = process.argv.slice(2);
curUser = Number(args[0]);
otherUser = 1 - curUser;
const socket = new WebSocket(`ws://localhost:5001/v1/ws/msg?token=${users[curUser].token}`);

socket.onopen = () => {
    console.log("Connected to server");

    const data = { user_id: users[otherUser].user_id, content: `Hi ${users[otherUser].user_id}` };
    socket.send(JSON.stringify(data));
};

socket.onmessage = (event) => {
    console.log("Received message from server:", event.data);

    setTimeout(() => {
        const data = { user_id: users[otherUser].user_id, content: `Hi ${users[otherUser].user_id}` };
        socket.send(JSON.stringify(data));
    }, 3000);
};

socket.onclose = () => {
    console.log("Connection closed");
};

socket.onerror = (error) => {
    console.log("Connection error", error)
}
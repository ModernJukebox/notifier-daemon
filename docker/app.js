const express = require("express");
const app = express();

app.use(express.json());

let counter = 0;
const notifications = new Map();

function getNextId() {
    counter += 1;

    return counter;
}

app.post('/clear', (req, res) => {
    notifications.clear();
    counter = 0;

    console.log("clear");

    return res.json("Cleared");
});

app.post("/notify", (req, res) => {
    const topics = req.query.topic
    const data = req.body;
    const id = getNextId();

    const notification = {
        date: new Date(),
        topics,
        data,
        id,
    };

    console.log("post", notification);

    notifications.set(id.toString(), notification);

    return res.status(201).json(notification);
});

app.listen(80, '0.0.0.0');

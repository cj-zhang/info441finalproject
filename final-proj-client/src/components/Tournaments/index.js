import React, { Component } from "react";
import "./style.css";


export default class Tournaments extends Component {
    postData(url, data) {
        // Default options are marked with *
        return fetch(url, {
            method: "POST", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, cors, *same-origin
            cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
            credentials: "same-origin", // include, *same-origin, omit
            headers: {
                "Content-Type": "application/json",
                // "Content-Type": "application/x-www-form-urlencoded",
            },
            redirect: "follow", // manual, *follow, error
            referrer: "no-referrer", // no-referrer, *client
            body: JSON.stringify(data), // body data type must match "Content-Type" header
        })
        .then(response => response.json()); // parses JSON response into native Javascript objects 
    }

    postData(`http://example.com/answer`, {answer: 42})
    .then(data => console.log(JSON.stringify(data))) // JSON-string from `response.json()` call
    .catch(error => console.error(error));

    

    render() {
        return (
            <h1>
                Tournament
            </h1>
        );
    }
}
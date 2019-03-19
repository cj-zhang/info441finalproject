import React, { Component } from "react";
import "./style.css";


export default class Tournaments extends Component {
    getTournaments(url) {
        // Default options are marked with *
        return fetch(url, {
            method: "GET", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, cors, *same-origin
            cache: "default", // *default, no-cache, reload, force-cache, only-if-cached
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

    

    

    render() {
        this.getTournaments("http://smash.chenjosephzhang.me/v1/tournaments")
        .then(data => console.log(JSON.stringify(data))) // JSON-string from `response.json()` call
        .catch(error => console.error(error));
        return (
            <h1>
                Tournament
            </h1>
        );
    }
}
import React, { Component } from "react";
import "./style.css";


export default class Tournaments extends Component {
    getTournaments(url) {
        // Default options are marked with *
        return fetch(url, {
            method: "GET", // *GET, POST, PUT, DELETE, etc.
            mode: "no-cors", // no-cors, cors, *same-origin
            cache: "default", // *default, no-cache, reload, force-cache, only-if-cached
            credentials: "same-origin", // include, *same-origin, omit
        })
        .then(response => response.json()); // parses JSON response into native Javascript objects 
    }

    

    

    render() {
        this.getTournaments("https://smash.chenjosephzhang.me/v1/tournaments")
        .then(data => console.log(JSON.stringify(data))) // JSON-string from `response.json()` call
        .catch(error => console.error(error));
        return (
            <h1>
                Tournament
            </h1>
        );
    }
}
import React, { Component } from "react";
import "./style.css";


export default class Tournaments extends Component {
    componentDidMount() {
        fetch("https://smash.chenjosephzhang.me/v1/tournaments", {
            mode: "no-cors",
        })
            .then(response => console.log(response))
            .then(data => console.log("hello:" + data))
            .catch(error => console.error(error));;
    }

    
    
    

    render() {
        let imgUrl = "https://smashgg.imgix.net/images/tournament/143919/image-b0e64c115be8f6274573d49974ff376f.jpg?auto=compress,format&w=300";
    
    
        let cardStyle = {
            width: "18rem",
        }
        return (
            <div>
                <h1 className="text-center mt-5">
                    Tournaments
                </h1>
                <div className="card mx-3" style={cardStyle}>
                    <img class="card-img-top" src="https://smashgg.imgix.net/images/tournament/143919/image-b0e64c115be8f6274573d49974ff376f.jpg?auto=compress,format&w=300" alt="Card image cap"/>
                    <div class="card-body">
                        <h5 class="card-title">Card title</h5>
                        <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
                        <a href="#" class="btn btn-primary">Go somewhere</a>
                    </div>
                </div>
            </div>
        );
    }
}
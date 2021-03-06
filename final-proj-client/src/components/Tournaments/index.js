import React, { Component } from "react";
import "./style.css";


export default class Tournaments extends Component {
    constructor(props) {
        super(props);
        this.state = {
            tourneyData: undefined,
        };
    }

    makeCard(jsonData) {
        for (let i = 0; i < jsonData.length; i++) {
            console.log(jsonData);
            console.log(jsonData[i].id);
            console.log(jsonData[i].url);
            console.log(jsonData[i].location);
            console.log(jsonData[i].organizer);
            console.log(jsonData[i].photoURL);
        }

    }

    componentDidMount() {
        fetch("https://smash.chenjosephzhang.me/v1/tournaments")
            .then(response => response.json())
            .then(data => this.makeCard(data))
            .catch(error => console.error(error));
    }

    render() {    
    
        let cardStyle = {
            height: 440,
            width: 350,
        }
        return (
            <div>
                <p className="text-center mt-5 header">
                    Featured Tournaments
                </p>
                <hr></hr>
                <div className="d-flex mb-5 justify-content-center"> 
                    <div className="card mx-4" style={cardStyle}>
                        <img className="card-img-top" src="https://smashgg.imgix.net/images/tournament/143919/image-b0e64c115be8f6274573d49974ff376f.jpg?auto=compress,format&w=300" alt="Card image cap"/>
                        <div className="card-body">
                            <h5 className="mb-1 card-title">Twitch Rivals: Showdown</h5>
                            <p className="mb-1 py-0 text-secondary">March 19th-21st 2019</p>
                            <p className="card-text">The Twitch Rivals: The Division 2 Showdown is an online competition, split between Europe and North America, featuring 40 invited streamers (20 per region). </p>
                            <a href="#" className="btn btn-primary mr-3">Details</a>
                            <a href="#" className="btn btn-success">Register</a>
                        </div>
                    </div>
                    <div className="card mx-4" style={cardStyle}>
                        <img className="card-img-top" src="https://smashgg.imgix.net/images/tournament/96984/image-9d8dd9be245d87fd520ec66cf5e4300d.png?auto=compress,format&w=300" alt="Card image cap"/>
                        <div className="card-body">
                            <h5 className="mb-1 card-title">Full Bloom 5</h5>
                            <p className="mb-1 py-0 text-secondary">March 23rd-24th 2019</p>
                            <p className="card-text">“Full Bloom 5” is a Melee/Smash Ultimate tournament held by Smash at IUB. Registration is mandatory and closes March 9th, 2019 at 11:59pm EDT. There will be no registration on site.</p>
                            <a href="#" className="btn btn-primary mr-3">Details</a>
                            <a href="#" className="btn btn-success">Register</a>
                        </div>
                    </div>
                    <div className="card mx-4" style={cardStyle}>
                        <img className="card-img-top" src="https://smashgg.imgix.net/images/tournament/128651/image-23f261b4cdf0edd3e858bcdec3e17b44.png?auto=compress,format&w=300" alt="Card image cap"/>
                        <div className="card-body">
                            <h5 className="mb-1 card-title">Smash Ultimate Summit</h5>
                            <p className="mb-1 py-0 text-secondary">March 8th-10th 2019</p>
                            <p className="card-text">Smash Summit is the most prestigious tournament and the sole invite-only event in the circuit. 20 of the world's best players will assemble to face off. </p>
                            <a href="#" className="btn btn-primary mr-3">Details</a>
                            <a href="#" className="btn btn-success">Register</a>
                        </div>
                    </div>
                </div>
                <div className="d-flex mb-5 justify-content-center"> 
                    <div className="card mx-4" style={cardStyle}>
                        <img className="card-img-top" src="https://smashgg.imgix.net/images/tournament/143919/image-b0e64c115be8f6274573d49974ff376f.jpg?auto=compress,format&w=300" alt="Card image cap"/>
                        <div className="card-body">
                            <h5 className="mb-1 card-title">Twitch Rivals: Showdown</h5>
                            <p className="mb-1 py-0 text-secondary">March 19th-21st 2019</p>
                            <p className="card-text">The Twitch Rivals: The Division 2 Showdown is an online competition, split between Europe and North America, featuring 40 invited streamers (20 per region). </p>
                            <a href="#" className="btn btn-primary mr-3">Details</a>
                            <a href="#" className="btn btn-success">Register</a>
                        </div>
                    </div>
                    <div className="card mx-4" style={cardStyle}>
                        <img className="card-img-top" src="https://smashgg.imgix.net/images/tournament/143919/image-b0e64c115be8f6274573d49974ff376f.jpg?auto=compress,format&w=300" alt="Card image cap"/>
                        <div className="card-body">
                            <h5 className="mb-1 card-title">Twitch Rivals: Showdown</h5>
                            <p className="mb-1 py-0 text-secondary">March 19th-21st 2019</p>
                            <p className="card-text">The Twitch Rivals: The Division 2 Showdown is an online competition, split between Europe and North America, featuring 40 invited streamers (20 per region). </p>
                            <a href="#" className="btn btn-primary mr-3">Details</a>
                            <a href="#" className="btn btn-success">Register</a>
                        </div>
                    </div>
                    <div className="card mx-4" style={cardStyle}>
                        <img className="card-img-top" src="https://smashgg.imgix.net/images/tournament/143919/image-b0e64c115be8f6274573d49974ff376f.jpg?auto=compress,format&w=300" alt="Card image cap"/>
                        <div className="card-body">
                            <h5 className="mb-1 card-title">Twitch Rivals: Showdown</h5>
                            <p className="mb-1 py-0 text-secondary">March 19th-21st 2019</p>
                            <p className="card-text">The Twitch Rivals: The Division 2 Showdown is an online competition, split between Europe and North America, featuring 40 invited streamers (20 per region). </p>
                            <a href="#" className="btn btn-primary mr-3">Details</a>
                            <a href="#" className="btn btn-success">Register</a>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}
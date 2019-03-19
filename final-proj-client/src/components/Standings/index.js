import React, { Component } from "react";
import "./style.css";


export default class Standings extends Component {
    render() {
        return (
            <div>
                <h1 className="text-center my-5">
                    Standings
                </h1>
                <div className="d-flex justify-content-around border my-1">
                    <div className="mx-5" id="container">
                        <div id="first"><img className="card-img-left mx-3" src="https://liquipedia.net/commons/images/1/19/Don%27t_Park_on_the_Grass_2018_logo.png" width="120" height="120" alt="Card image cap"/></div>
                        <div id="second">
                            <h4 className="py-1">Don't Park On the Grass</h4> <br/>
                            <p className="mb-1 py-0 text-secondary">December 19th-21st 2018</p>
                        </div>
                    </div>
                    <div>
                            <p className="px-3 py-1"><span className="text-success">WON 2-0</span> against <strong>Jo96</strong></p>
                            <p className="px-3 py-1"><span className="text-success">WON 2-1</span> against <strong>Sunwoowoo</strong></p>
                            <p className="px-3 py-1"><span className="text-danger">LOST 1-2</span> against <strong>KxuOJOM</strong></p>

                    </div>
                </div>
                <div className="d-flex justify-content-around border my-1">
                    <div className="mx-5" id="container">
                        <div id="first"><img className="card-img-left mx-3" src="https://upload.wikimedia.org/wikipedia/en/thumb/3/36/Evo_Championship_Series_Logo.png/200px-Evo_Championship_Series_Logo.png" width="120" height="120" alt="Card image cap"/></div>
                        <div id="second">
                            <h4 className="py-1">EVOLUTION 2018</h4> <br/>
                            <p className="mb-1 py-0 text-secondary">July 23rd-28th 2018</p>
                        </div>
                    </div>
                    <div>
                            <p className="px-3 py-1"><span className="text-success">WON 2-1</span> against <strong>Hbox101</strong></p>
                            <p className="px-3 py-1"><span className="text-danger">LOST 0-2</span> against <strong>DowellInfo</strong></p>

                    </div>
                </div>
                <div className="d-flex justify-content-around border my-1">
                    <div className="mx-5" id="container">
                        <div id="first"><img className="card-img-left mx-3" src="https://upload.wikimedia.org/wikipedia/en/thumb/4/4f/The_logo_of_The_Big_House_Six.png/220px-The_logo_of_The_Big_House_Six.png" width="120" height="120" alt="Card image cap"/></div>
                        <div id="second">
                            <h4 className="py-1">The Big House 8</h4> <br/>
                            <p className="mb-1 py-0 text-secondary">October 10th-13th 2018</p>
                        </div>
                    </div>
                    <div>
                            <p className="px-3 py-1"><span className="text-success">WON 2-0</span> against <strong>Jo96</strong></p>
                            <p className="px-3 py-1"><span className="text-success">WON 2-1</span> against <strong>Sunwoowoo</strong></p>
                            <p className="px-3 py-1"><span className="text-danger">LOST 1-2</span> against <strong>KxuOJOM</strong></p>

                    </div>
                </div>
               
            </div>
        );
    }
}
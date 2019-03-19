import React, { Component } from "react";
import "./style.css";


export default class Standings extends Component {
    render() {
        return (
            <div>
                <h1 className="text-center my-5">
                    Standings
                </h1>
                <div className="card">
                    <div className="card-body">
                        <h5 className="card-title">Don't Park On the Grass</h5>
                        <p className="mb-1 py-0 text-secondary">December 19th-21st 2018</p>
                        <p className="card-text"><span className="text-success">WON 2-0</span> against <strong>Jo96</strong></p>
                    </div>
                </div>
            </div>
        );
    }
}
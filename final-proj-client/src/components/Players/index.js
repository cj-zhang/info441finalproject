import React, { Component } from "react";
import "./style.css";

export default class Contact extends Component {
    render() {
        let cardStyle = {
            height: 170,
            width: 600,
            marginTop: 10,
            marginBottom: 15,
        }
        return (
            <div class="container">
                <div class="row">
                    <div class="col-md-12">
                        <div class="card">
                                <p className="text-center mt-5 header" >
                                    Players</p>
                            <hr></hr>
                            <div class="gaadiex-list">
                                <div class="gaadiex-list-item-text">
                                    <div className="card mx-4" style={cardStyle}>
                                        <div className="card-body">
                                            <h5 className="mb-1 card-title">Kevxuxu</h5>
                                            <p className="mb-1 py-0 text-secondary">Latest Tournament: UW Locals</p>
                                            <p>Placement: 1st Place</p>    
                                            <a href="#" className="btn btn-primary mr-3">Details</a>                                    
                                        </div>
                                    </div>
                                </div>
                                <div class="gaadiex-list-item-text">
                                <div className="card mx-4" style={cardStyle}>
                                        <div className="card-body">
                                            <h5 className="mb-1 card-title">Jozhazha</h5>
                                            <p className="mb-1 py-0 text-secondary">Latest Tournament: UW Locals</p>
                                            <p>Placement: Last Place</p>    
                                            <a href="#" className="btn btn-primary mr-3">Details</a>                                    
                                        </div>
                                    </div>
                                </div>
                                <div className="card mx-4" style={cardStyle}>
                                        <div className="card-body">
                                            <h5 className="mb-1 card-title">Sunwoowoo</h5>
                                            <p className="mb-1 py-0 text-secondary">Latest Tournament: UW Locals</p>
                                            <p>Placement: 2nd Place</p>    
                                            <a href="#" className="btn btn-primary mr-3">Details</a>                                    
                                        </div>
                                    </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}
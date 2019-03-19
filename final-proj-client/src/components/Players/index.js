import React, { Component } from "react";
import "./style.css";

export default class Contact extends Component {
    render() {
        return (
            <div class="container">
                <div class="row">
                    <div class="col-md-12">
                        <div class="card">
                            <div class="row">
                                <h2 class="col-md-8 card-header">Players</h2>
                                <input class="col-md-4" id="srchbar" type="search" placeholder="Search" />
                            </div>
                            <div class="gaadiex-list">
                                <div class="gaadiex-list-item"><img class="gaadiex-list-item-img"
                                    src="/Users/sunwookang/Downloads/character_renders/Wolf.png" alt="List user" />
                                    <div class="gaadiex-list-item-text">
                                        <h3><a href="#">Kevxuxu</a></h3>
                                        <h4>Latest Tournament: UW Locals</h4>
                                        <p>Placement: 1st Place</p>
                                    </div>
                                </div>
                                <div class="gaadiex-list-item"><img class="gaadiex-list-item-img"
                                    src="/Users/sunwookang/Downloads/character_renders/Lucina.png" alt="List user" />
                                    <div class="gaadiex-list-item-text">
                                        <h3><a href="#">Jozhazha</a></h3>
                                        <h4>Latest Tournament: UW Locals</h4>
                                        <p>Placement: Last Place</p>
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
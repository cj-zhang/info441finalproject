import React, { Component } from "react";
import jumbotron from "react-bootstrap"
import { Row, Col, Container } from 'react-bootstrap';
import "./style.css";
import { Link } from "react-router-dom";
import Tournaments from "../Tournaments";


export default class Home extends Component {
    render() {
        return (
            <div>
                <Row className="mainjumbotron">
                    <Col className="container title-container">
                        <h1 className="smash-title">Smash.qq</h1>
                        <h2 className="title-text">Find your local Super Smash Bros. Ultimate tournaments</h2>
                        <h3 className="signuplink">Sign Up <Link to={"/login"} className="now">Now</Link>!</h3>
                    </Col>
                </Row>
                <Tournaments/>
                <br></br>
                <div id="root"></div>
            </div>
        );
    }
}
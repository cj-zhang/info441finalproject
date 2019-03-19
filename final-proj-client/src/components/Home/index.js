import React, { Component } from "react";
import jumbotron from "react-bootstrap"
import { Row, Col, Container } from 'react-bootstrap';
import "./style.css";

export default class Home extends Component {
    render() {
        return (
            <div>
                <Row className="jumbotron">
                    <Col className="container title-container">
                        <h1 className="smash-title">Smash.qq</h1>
                        <h2 className="title-text">Find your local Super Smash Bros. Ultimate tournaments</h2>
                        <h3 className="signuplink">Sign Up <a className="now"
                            href="file:///Users/sunwookang/go/src/info441finalproject/final-proj-client/public/pages/signup/signup.html">Now</a>!</h3>
                    </Col>
                </Row>
                <div className="container-fluid bg-3 text-center">
                    <h3>Local tournaments</h3>
                    <Row className="row">
                        <Col sm={3} className="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" className="img-responsive" style={{width: '100%'}} alt="Image"></img>
                        </Col>
                        <Col sm={3} className="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" className="img-responsive" style={{width: '100%'}} alt="Image"></img>
                        </Col>
                        <Col sm={3} className="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" className="img-responsive" style={{width: '100%'}} alt="Image"></img>
                        </Col>
                        <Col sm={3} className="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" className="img-responsive" style={{width: '100%'}} alt="Image"></img>
                        </Col>

                    </Row>
                </div>

                <div class="container-fluid bg-3 text-center">
                    <div class="row">
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style={{width: '100%'}} alt="Image"></img>
                        </div>
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style={{width: '100%'}} alt="Image"></img>
                        </div>
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style={{width: '100%'}} alt="Image"></img>
                        </div>
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style={{width: '100%'}} alt="Image"></img>
                        </div>
                    </div>
                </div>
                <br></br>
                <div id="root"></div>
            </div>
        );
    }
}
import React, { Component } from "react";
import "./style.css";

export default class Home extends Component {
    render() {
        return (
            <div>
                <div class="jumbotron">
                    <div class="container title-container">
                        <h1 class="smash-title">Smash.qq</h1>
                        <h2 class="title-text">Find your local Super Smash Bros. Ultimate tournaments</h2>
                        <h3 class="signuplink">Sign Up <a class="now"
                            href="file:///Users/sunwookang/go/src/info441finalproject/final-proj-client/public/pages/signup/signup.html">Now</a>!</h3>
                    </div>
                </div>
                <div class="container-fluid bg-3 text-center">
                    <h3>Local tournaments</h3>
                    <div class="row">
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style="width:100%" alt="Image"></img>
                        </div>
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style="width:100%" alt="Image"></img>
                        </div>
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style="width:100%" alt="Image"></img>
                        </div>
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style="width:100%" alt="Image"></img>
                        </div>

                    </div>
                </div>

                <div class="container-fluid bg-3 text-center">
                    <div class="row">
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style="width:100%" alt="Image"></img>
                        </div>
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style="width:100%" alt="Image"></img>
                        </div>
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style="width:100%" alt="Image"></img>
                        </div>
                        <div class="col-sm-3">
                            <p>Some text..</p>
                            <img src="https://placehold.it/150x80?text=IMAGE" class="img-responsive" style="width:100%" alt="Image"></img>
                        </div>
                    </div>
                </div>
                <br></br>
                <div id="root"></div>
            </div>
        );
    }
}
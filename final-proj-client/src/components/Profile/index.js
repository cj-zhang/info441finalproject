import React, { Component } from "react";
import "/style.css";

export default class Profile extends Component {
    render() {
        return (
            <div>
                <div class="container emp-profile">
                    <form method="post">
                        <div class="row">
                            <div class="col-md-4">
                                <div class="profile-img">
                                    <img src="/Users/sunwookang/Downloads/character_renders/KingDedede.png" alt="" />
                                    <div class="file btn btn-lg btn-primary">Change Character</div>
                                </div>
                            </div>
                            <div class="col-md-6">
                                <div class="profile-head">
                                    <h2>User Tag</h2>
                                    <p class="profile-rating">Most Recent Placement <span>#3</span></p>
                                    <ul class="nav nav-tabs" id="myTab" role="tablist">
                                        <li class="nav-item">
                                            <a class="nav-link" id="home-tab" data-toggle="tab" href="#home" role="tab" aria-controls="home"
                                                aria-selected="true">About</a>
                                        </li>
                                        <li class="nav-item">
                                            <a class="nav-link" id="profile-tab" data-toggle="tab" href="#profile" role="tab"
                                                aria-controls="profile" aria-selected="false">Previous Tournaments</a>
                                        </li>
                                    </ul>
                                </div>
                            </div>
                            <div class="col-md-2">
                                <input type="submit" class="profile-edit-btn" name="btnAddMore" value="Edit Profile" />
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-md-4">
                                <div class="profile-work">
                                    <p>Other Characters Played</p>
                                    <img src="/Users/sunwookang/Downloads/character_renders/Mario.png" alt="" class="otherchar" />
                                    <img src="/Users/sunwookang/Downloads/character_renders/DonkeyKong.png" alt="" class="otherchar" />
                                    <img src="/Users/sunwookang/Downloads/character_renders/Link.png" alt="" class="otherchar" />
                                </div>
                            </div>
                            <div class="col-md-8">
                                <div class="tab-content profile-tab" id="myTabContent">
                                    <div class="tab-pane active" id="home" role="tabpanel" aria-labelledby="home-tab">
                                        <div class="row">
                                            <div class="col-md-6">
                                                <label>User Tag</label>
                                            </div>
                                            <div class="col-md-6">
                                                <p>Sunwoowoo</p>
                                            </div>
                                        </div>
                                        <div class="row">
                                            <div class="col-md-6">
                                                <label>Name</label>
                                            </div>
                                            <div class="col-md-6">
                                                <p>Sunny Kang</p>
                                            </div>
                                        </div>
                                        <div class="row">
                                            <div class="col-md-6">
                                                <label>Email</label>
                                            </div>
                                            <div class="col-md-6">
                                                <p>helpme@gmail.com</p>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="tab-pane fade" id="profile" role="tabpanel" aria-labelledby="profile-tab">
                                        <div class="row">
                                            <div class="col-md-4">
                                                <label>Tournament</label>
                                            </div>
                                            <div class="col-md-4">
                                                <label>Record</label>
                                            </div>
                                            <div class="col-md-4">
                                                <label>Placing</label>
                                            </div>
                                        </div>
                                        <div class="row">
                                            <div class="col-md-4">
                                                <p>Genesis 6</p>
                                            </div>
                                            <div class="col-md-4">
                                                <p>0-2</p>
                                            </div>
                                            <div class="col-md-4">
                                                <p>Did not exit pools</p>
                                            </div>
                                        </div>
                                        <div class="row">
                                            <div class="col-md-4">
                                                <p>Frostbite</p>
                                            </div>
                                            <div class="col-md-4">
                                                <p>0-2</p>
                                            </div>
                                            <div class="col-md-4">
                                                <p>Did not exit pools</p>
                                            </div>
                                        </div>
                                        <div class="row">
                                            <div class="col-md-4">
                                                <p>Summit</p>
                                            </div>
                                            <div class="col-md-4">
                                                <p>0-2</p>
                                            </div>
                                            <div class="col-md-4">
                                                <p>Did not exit pools</p>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        );
    }
}

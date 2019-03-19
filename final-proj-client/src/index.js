import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Home from "./components/Home";
import Profile from "./components/Profile";
import Players from "./components/Players";
import Login from "./components/Login";
import Signup from "./components/Signup";
import Games from "./components/Games";
import Tournaments from "./components/Tournaments";
import Standings from "./components/Standings";
import "./index.css";
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import Permanav from "./components/Layout/Header";


ReactDOM.render(
    <Router>
        <Permanav />
        
        <Route exact path="/" component={Home} />
        <Route path="/profile" component={Profile} />
        <Route path="/players" component={Players} />
        <Route path="/login" component={Login} />
        <Route path="/signup" component={Signup} />
        {/* <Route path="/games" component={Games} /> */}
        <Route path="/tournaments" component={Tournaments} />
        <Route path="/standings" component={Standings} />

    </Router>,
    document.getElementById("root")
);

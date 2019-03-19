import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Layout from "./components/Layout";
import Home from "./components/Home";
import Profile from "./components/Profile";
import Players from "./components/Players";
import Login from "./components/Login";
import Signup from "./components/Signup";
import "./index.css";
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';

ReactDOM.render(
    <Router>
        {/* <Route path="/" component={Layout} /> */}
        <Route exact path="/" component={Home} />
        <Route path="/profile" component={Profile} />
        <Route path="/players" component={Players} />
        <Route path="/login" component={Login} />
        <Route path="/signup" component={Signup} />

    </Router>,
    document.getElementById("root")
);

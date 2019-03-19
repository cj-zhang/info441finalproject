import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Layout from "./components/Layout";
import Home from "./components/Home";
import Profile from "./components/Profile";
import Players from "./components/Players";
import "./index.css";

const outlet = document.getElementById("root");
ReactDOM.render(
    <Router exact path="/" component={Layout}>
        <Route component={Home} />
        <Route path="/profile" component={Profile} />
        <Route path="/players" component={Players} />
    </Router>,
    document.getElementById("root")
);

import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Layout from "./components/Layout";
import Home from "./components/Home";
import Profile from "./components/Profile";
import Players from "./components/Players";

const Routes = (
  <Router path="/" component={Layout}>
    <Route component={Home} />
    <Route path="/profile" component={Profile} />
    <Route path="/players" component={Players} />
  </Router>
);

export default Routes;

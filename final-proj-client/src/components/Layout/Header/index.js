import React, { Component } from "react";
import { Nav, Navbar } from 'react-bootstrap';
import "./style.css";

export default class Permanav extends Component {
	render() {
		return (
				<div className="header-div">
					<nav className="navbar navbar-expand-lg navbar-dark bg-dark">
						<a className="navbar-brand" href="/">Smash.qq</a>
						<button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
							<span className="navbar-toggler-icon"></span>
						</button>

						<div className="collapse navbar-collapse" id="navbarSupportedContent">
							<ul className="navbar-nav mr-auto">
								<li className="nav-item active">
									<a className="nav-link" href="/">Home <span className="sr-only">(current)</span></a>
								</li>
								<li className="nav-item">
									<a className="nav-link" href="/players">Players</a>
								</li>
								<li className="nav-item dropdown">
									<a className="nav-link" href="/tournaments">Tournaments</a>
								</li>
								<li className="nav-item">
									<a className="nav-link" href="/games">Games</a>
								</li>
								<li className="nav-item">
									<a className="nav-link" href="/standings">Standings</a>
								</li>
							</ul>
							<a className="nav-link" href="/profile">My Profile</a>
							<a className="nav-link" href="/login">Sign In</a>
						</div>
					</nav>
				</div>

				);
			}
		}

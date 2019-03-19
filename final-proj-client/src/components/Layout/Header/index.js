import React, { Component } from "react";
import {Nav} from 'react-bootstrap';
import {Navbar} from 'react-bootstrap';
import "./style.css";

export default class Header extends Component {
	render() {
		return (
			<div className="header-div">
				<Navbar fixed="top">
					<Nav className="justify-content-end" activeKey="/home">
					<Nav.Item>
					<Nav.Link href="/">Home</Nav.Link>
					</Nav.Item>
					<Nav.Item>
					<Nav.Link href="/players">Players</Nav.Link>
					</Nav.Item>
					<Nav.Item>
					<Nav.Link href="/tournaments">Tournaments</Nav.Link>
					</Nav.Item>
                    <Nav.Item>
					<Nav.Link href="/games">Games</Nav.Link>
					</Nav.Item>
                    <Nav.Item>
					<Nav.Link href="/standings">Standings</Nav.Link>
					</Nav.Item>
					</Nav>
				</Navbar>
			</div>
		);
	}
}

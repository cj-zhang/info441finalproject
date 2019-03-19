import React from "react";
import { Link } from "react-router-dom";


export default class SignUpView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            errorMessage: undefined,
            currentUser: undefined,
            email: "",
            password: "",
            confirm: "",
            displayName: "",
            firstName: "",
            lastName: "",
            photoURL: ""
        };
    }
    // componentDidMount() {
    //     this.authUnsub = firebase.auth().onAuthStateChanged(user => {
    //         this.setState({currentUser: user});
    //     });
    // }
    // componentWillUnmount() {
    //     this.authUnsub();
    // }
    handleSignUp(evt) {
        var data = {
            email: this.state.email,
            password: this.state.password,
            passwordConf: this.state.confirm,
            userName: this.state.displayName,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
        }
        if (this.state.password !== this.state.confirm) {
            alert("Your passwords do not match");
        } 
        evt.preventDefault();
        return fetch('https://smash.chenjosephzhang.me/v1/users', {
            method: "POST", 
            mode: "cors", // no-cors, cors, *same-origin
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data), // body data type must match "Content-Type" header
        })
        .then(function(response) {
            console.log(response);
            if(response.status === 201) {
                console.log("SUCCESSFUL");
                for(let header of response.headers){
                    console.log("header: " + header);
                 }
            } else {
                alert("Unsuccessful sign up attempt");
            }
        })
        .then(() => this.props.history.push("/login"))
        .catch(function(error) {
            console.log('There has been a problem with your fetch operation: ', error.message);
        });
    }
            
    render() {
        let signupStyle = {
            width: "30%",
            marginTop: "60px",
            marginLeft: "auto",
            marginRight: "auto",
            backgroundColor: 'rgba(0,0,0,0.2)',
            borderRadius: '8px'
        }
        let titleStyle = {
            textAlign: "center",
            fontSize: 60,
            color: "white"
        }
        let textStyle = {
            marginLeft: "8px",
            paddingTop: "5px"
        }
        let labelStyle = {
            color: "white"
        }
        return (
            <div className="container" style={signupStyle}>
                <h1 style={titleStyle}>Sign Up</h1>

                <form onSubmit={evt => this.handleSignUp(evt)}>
                    <div className="form-group">
                        <h4 style={labelStyle}>Email:</h4>
                        <input id="email" type="email" className="form-control"
                            placeholder="enter your email address"
                            onInput={evt => this.setState({ email: evt.target.value })} />
                    </div>
                    <div className="form-group">
                        <h4 style={labelStyle}>Password:</h4>
                        <input id="password" type="password" className="form-control"
                            placeholder="enter your password"
                            onInput={evt => this.setState({ password: evt.target.value })} />
                    </div>
                    <div className="form-group">
                        <h4 style={labelStyle}>Confirm password:</h4>
                        <input id="confirm-password" type="password" className="form-control"
                            placeholder="confirm password"
                            onInput={evt => this.setState({ confirm: evt.target.value })} />
                    </div>
                    <div className="form-group">
                        <h4 style={labelStyle}>UserName:</h4>
                        <input id="display-name" type="display-name" className="form-control"
                            placeholder="enter your display name"
                            onInput={evt => this.setState({ displayName: evt.target.value })} />
                    </div>
                    <div className="form-group">
                        <h4 style={labelStyle}> First Name:</h4>
                        <input id="first-name" type="first-name" className="form-control"
                            placeholder="enter your first name"
                            onInput={evt => this.setState({ firstName: evt.target.value })} />
                    </div>
                    <div className="form-group">
                        <h4 style={labelStyle}>Last Name:</h4>
                        <input id="last-name" type="last-name" className="form-control"
                            placeholder="enter your last name"
                            onInput={evt => this.setState({ lastName: evt.target.value })} />
                    </div>
                    <div className="last-row d-flex">
                        <div className="form-group">
                            <button type="submit" className="btn btn-primary">
                                Sign Up
                            </button>
                        </div>
                        <p style={textStyle}>Already have an account? <Link to={"/login"}>Sign In</Link></p>
                    </div>
                </form>
            </div>
        );
    }
}
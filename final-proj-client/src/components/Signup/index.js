import React from "react";
import { Link } from "react-router-dom";


export default class SignUpView extends React.Component {
    // constructor(props) {
    //     super(props);
    //     this.state = {
    //         errorMessage: undefined,
    //         currentUser: undefined,
    //         email: "",
    //         password: "",
    //         confirm: "",
    //         displayName: "",
    //         photoURL: ""
    //     };
    // }
    // componentDidMount() {
    //     this.authUnsub = firebase.auth().onAuthStateChanged(user => {
    //         this.setState({currentUser: user});
    //     });
    // }
    // componentWillUnmount() {
    //     this.authUnsub();
    // }
    // handleSignUp(evt) {
    //     if (this.state.password !== this.state.confirm) {
    //         alert("Your passwords do not match");
    //     }
    //     evt.preventDefault();
    //     if (this.state.displayName) {
    //         this.setState({photoURL: "https://www.gravatar.com/avatar/" + md5(this.state.email.toLowerCase().trim()) + "?s=30"});
    //         firebase.auth().createUserWithEmailAndPassword(this.state.email, this.state.password)
    //             .then(user => {
    //                 user.updateProfile({
    //                     displayName: this.state.displayName,
    //                     photoURL: this.state.photoURL
    //                 })
    //             })
    //             .then(() => firebase.auth().signInWithEmailAndPassword(this.state.email, this.state.password))
    //             .then(() => this.props.history.push("/channels/General"))
    //             .catch(err => alert(err.message));
    //     } else {
    //         alert("Need a display name");
    //     }
    // }
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

                <form>
                    <div className="form-group">
                        <h4 style={labelStyle}>Email:</h4>
                        <input id="email" type="email" className="form-control"
                            placeholder="enter your email address"/>
                    </div>
                    <div className="form-group">
                        <h4 style={labelStyle}>Password:</h4>
                        <input id="password" type="password" className="form-control"
                            placeholder="enter your password"/>
                    </div>
                    <div className="form-group">
                        <h4 style={labelStyle}>Confirm password:</h4>
                        <input id="confirm-password" type="password" className="form-control"
                            placeholder="confirm password"/>
                    </div>
                    <div className="form-group">
                        <h4 style={labelStyle}>Display name:</h4>
                        <input id="display-name" type="display-name" className="form-control"
                            placeholder="enter your display name"/>
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
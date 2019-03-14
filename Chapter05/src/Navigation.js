import React from 'react';
import { NavLink } from 'react-router-dom';


export default class Navigation extends React.Component {

    buildLoggedInMenu() {
        return (
            <div className="navbar-brand order-1 text-white my-auto">
                <div className="btn-group">
                    <button type="button" className="btn btn-success dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                        Welcome {this.props.user.name}
                    </button>
                    <div className="dropdown-menu">
                        <a className="btn dropdown-item" role="button">Sign Out</a>
                    </div>
                </div>
            </div>
        );
    }

    render() {
        return (
            <div>
                <nav className="navbar navbar-expand-lg navbar-dark bg-success fixed-top">
                    <div className="container">
                        {
                            this.props.user.loggedin ?
                                /*<p className="navbar-brand order-1 text-white my-auto">Welcome {this.props.user.name}</p>*/
                                this.buildLoggedInMenu()
                                : <button type="button" className="navbar-brand order-1 btn btn-success" data-toggle="modal" data-target="#register">Sign in</button>
                        }
                        <div className="navbar-collapse" id="navbarNavAltMarkup">
                            <div className="navbar-nav">
                                <NavLink className="nav-item nav-link" to="/">Home</NavLink>
                                <NavLink className="nav-item nav-link" to="/promos">Promotions</NavLink>
                                {this.props.user.loggedin ? <NavLink className="nav-item nav-link" to="/myorders">My Orders</NavLink> : null}
                                <NavLink className="nav-item nav-link" to="/about">About</NavLink>
                            </div>
                        </div>
                    </div>
                </nav>
            </div>
        );
    }
}
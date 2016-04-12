"use strict";

import React from 'react';
import ReactDOM from 'react-dom';
import {Router, Route, Link, browserHistory, IndexRoute} from 'react-router';
import AllPosts from './all-posts.jsx';
import SinglePost from './single-post.jsx';
import NewPost from './new-post.jsx';
import stupidThingDoesGlobals from 'whatwg-fetch';


class Header extends React.Component {
  constructor(props){
    super(props);
  }
  
  render(){
    return (
      <div>
        <h1>Welcome To Blog</h1>
        <ul>
          <li><Link to="/">Home</Link></li>
          <li><Link to="/new-post">New Post</Link></li>
        </ul>
        {this.props.children}
      </div>
    );
  }
  
}

class App extends React.Component {
  constructor(props){
    super(props);
  }
  
  render(){
    return (
      <Router history={browserHistory}>
        <Route path="/" component={Header}>
          <IndexRoute component={AllPosts} />
          <Route path="post/:postId" component={SinglePost} />
          <Route path="new-post" component={NewPost} />
        </Route>
      </Router>
    );
  }
}

ReactDOM.render(<App/>, document.getElementById('container'));

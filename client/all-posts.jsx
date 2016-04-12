"use strict";
import React from 'react';
import {Link} from 'react-router';

const PostsList = props => {
  const list = props.posts.map(x => <li><Link to={'/post/' + x.id}><h3>{x.title}</h3></Link><p>{x.body}</p></li>);
  return <ul>{list}</ul>;
};

export default class AllPosts extends React.Component{
  constructor(props){
    super(props);
    this.state = {
      posts : []
    };
  }
  
  componentWillMount(){
    this.getPosts()
    .then(posts => this.setState({posts}))
    .catch(err => console.log(err));
  }
  
  getPosts(){
    return fetch('/api/posts')
    .then(response => response.json())
    .then(posts => posts.map( (x, i) => Object.assign(x, {id : i}) ));
  }
  
  render(){
    return (
      <div>
        <h1>Welcome to my blog</h1>
        <PostsList posts={this.state.posts} />
        {this.props.children}
      </div>
    );
  }
  
}



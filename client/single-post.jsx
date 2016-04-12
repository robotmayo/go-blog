"use strict";

import React from 'react';

export default class SinglePost extends React.Component {
  constructor(props){
    super(props);
    this.state = {
      post : {
        id : 0,
        title : '',
        body : ''
      }
    };
  }
  
  componentWillMount(){
    this.getPost(this.props.params.postId)
    .then(post => this.setState({post}))
    .catch(err => console.log(err));
  }
  
  getPost(id){
    return fetch(`/api/post/${id}`)
    .then(response => response.json())
    .then(post => {
      console.log(post);
      return post;
    })
  }
  
  render(){
    return (
      <div>
        <h1>{this.state.post.title}</h1>
        <p>{this.state.post.body}</p>
      </div>
    );
  }
  
}
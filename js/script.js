// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: May 2023
// This file contains the JS functions for index.html

'use strict'

function myButtonClicked() {
  // This function tells he user if they get free admission
  // input
  const userAge = parseInt(document.getElementById("user-age").value)
  const dayOfTheWeek = document.getElementById("day-of-the-week").value

  // process
  if ((dayOfTheWeek == "Monday" || dayOfTheWeek == "Thursday") || (userAge > 12 && userAge < 21)) {
    // output
    document.getElementById("answer").innerHTML = "You get free admission!"
  } else {
    // output
    document.getElementById("answer").innerHTML = "You have to pay admission."
  }
}

// Copyright (c) 2023 Alyssia fung All rights reserved
//
// Created by: Alyssia fung
// Created on: Mar 2023
// This file contains the JS functions for index.html

/**
 * This function calculates area of a circle
 */
"use strict"

function calculate() {
  // input
  const radius = parseInt(document.getElementById("radius").value)

  // process
  const area = 3.14 * radius ** 2

  // output
  document.getElementById("area").innerHTML = "Area is: " + area + " cm²"
}

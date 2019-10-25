package bund_context

import (
  "github.com/common-nighthawk/go-figure"
  "github.com/hashicorp/go-uuid"
)

var Version = "MonitoringBund version 0.1.1 (development release)"
var Version_Num = "0.1.1"
var Version_Release = "(development release)"
var Logo = figure.NewFigure("[theBund> ", "o8", true)
var PS1 = "[BUND> "

var (
  CfgFile     string
  Logverbose  string
)

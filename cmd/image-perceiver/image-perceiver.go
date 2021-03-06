/*
Copyright (C) 2018 Synopsys, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"fmt"
	"os"

	"github.com/blackducksoftware/perceivers/cmd/image-perceiver/app"
	"github.com/blackducksoftware/perceivers/pkg/annotations"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting image-perceiver")
	configPath := os.Args[1]
	log.Printf("Config path: %s", configPath)
	handler := annotations.ImageAnnotatorHandlerFuncs{
		ImageLabelCreationFunc:      annotations.CreateImageLabels,
		ImageAnnotationCreationFunc: annotations.CreateImageAnnotations,
		MapCompareHandlerFuncs: annotations.MapCompareHandlerFuncs{
			MapCompareFunc: annotations.StringMapContains,
		},
	}
	// Create the Image Perceiver
	perceiver, err := app.NewImagePerceiver(handler, configPath)
	if err != nil {
		panic(fmt.Errorf("failed to create image-perceiver: %v", err))
	}

	// Run the perceiver
	stopCh := make(chan struct{})
	perceiver.Run(stopCh)
}

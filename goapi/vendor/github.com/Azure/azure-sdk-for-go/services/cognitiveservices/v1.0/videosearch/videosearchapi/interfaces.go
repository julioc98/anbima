package videosearchapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/videosearch"
)

// VideosClientAPI contains the set of methods on the VideosClient type.
type VideosClientAPI interface {
	Details(ctx context.Context, query string, acceptLanguage string, userAgent string, clientID string, clientIP string, location string, countryCode string, ID string, modules []videosearch.VideoInsightModule, market string, resolution videosearch.VideoResolution, safeSearch videosearch.SafeSearch, setLang string, textDecorations *bool, textFormat videosearch.TextFormat) (result videosearch.VideoDetails, err error)
	Search(ctx context.Context, query string, acceptLanguage string, userAgent string, clientID string, clientIP string, location string, countryCode string, count *int32, freshness videosearch.Freshness, ID string, length videosearch.VideoLength, market string, offset *int32, pricing videosearch.VideoPricing, resolution videosearch.VideoResolution, safeSearch videosearch.SafeSearch, setLang string, textDecorations *bool, textFormat videosearch.TextFormat) (result videosearch.Videos, err error)
	Trending(ctx context.Context, acceptLanguage string, userAgent string, clientID string, clientIP string, location string, countryCode string, market string, safeSearch videosearch.SafeSearch, setLang string, textDecorations *bool, textFormat videosearch.TextFormat) (result videosearch.TrendingVideos, err error)
}

var _ VideosClientAPI = (*videosearch.VideosClient)(nil)

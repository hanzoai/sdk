// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsKeyListResponseKeysUserAPIKeyAuthExpiresUnion() {}
func (UnionTime) ImplementsKeyBlockResponseExpiresUnion()                  {}
func (UnionTime) ImplementsSpendListLogsResponseEndTimeUnion()             {}
func (UnionTime) ImplementsSpendListLogsResponseStartTimeUnion()           {}
func (UnionTime) ImplementsSpendListTagsResponseEndTimeUnion()             {}
func (UnionTime) ImplementsSpendListTagsResponseStartTimeUnion()           {}
func (UnionTime) ImplementsGlobalSpendListTagsResponseEndTimeUnion()       {}
func (UnionTime) ImplementsGlobalSpendListTagsResponseStartTimeUnion()     {}
func (UnionTime) ImplementsGlobalSpendGetReportResponseEndTimeUnion()      {}
func (UnionTime) ImplementsGlobalSpendGetReportResponseStartTimeUnion()    {}

type UnionString string

func (UnionString) ImplementsModelNewParamsLlmParamsConfigurableClientsideAuthParamUnion()         {}
func (UnionString) ImplementsModelNewParamsLlmParamsStreamTimeoutUnion()                           {}
func (UnionString) ImplementsModelNewParamsLlmParamsTimeoutUnion()                                 {}
func (UnionString) ImplementsUpdateDeploymentLlmParamsConfigurableClientsideAuthParamsUnionParam() {}
func (UnionString) ImplementsUpdateDeploymentLlmParamsStreamTimeoutUnionParam()                    {}
func (UnionString) ImplementsUpdateDeploymentLlmParamsTimeoutUnionParam()                          {}
func (UnionString) ImplementsFineTuningJobNewParamsHyperparametersBatchSizeUnion()                 {}
func (UnionString) ImplementsFineTuningJobNewParamsHyperparametersLearningRateMultiplierUnion()    {}
func (UnionString) ImplementsFineTuningJobNewParamsHyperparametersNEpochsUnion()                   {}
func (UnionString) ImplementsKeyListResponseKeysUnion()                                            {}
func (UnionString) ImplementsKeyListResponseKeysUserAPIKeyAuthExpiresUnion()                       {}
func (UnionString) ImplementsKeyBlockResponseExpiresUnion()                                        {}
func (UnionString) ImplementsSpendListLogsResponseEndTimeUnion()                                   {}
func (UnionString) ImplementsSpendListLogsResponseMessagesUnion()                                  {}
func (UnionString) ImplementsSpendListLogsResponseResponseUnion()                                  {}
func (UnionString) ImplementsSpendListLogsResponseStartTimeUnion()                                 {}
func (UnionString) ImplementsSpendListTagsResponseEndTimeUnion()                                   {}
func (UnionString) ImplementsSpendListTagsResponseMessagesUnion()                                  {}
func (UnionString) ImplementsSpendListTagsResponseResponseUnion()                                  {}
func (UnionString) ImplementsSpendListTagsResponseStartTimeUnion()                                 {}
func (UnionString) ImplementsGlobalSpendListTagsResponseEndTimeUnion()                             {}
func (UnionString) ImplementsGlobalSpendListTagsResponseMessagesUnion()                            {}
func (UnionString) ImplementsGlobalSpendListTagsResponseResponseUnion()                            {}
func (UnionString) ImplementsGlobalSpendListTagsResponseStartTimeUnion()                           {}
func (UnionString) ImplementsGlobalSpendGetReportResponseEndTimeUnion()                            {}
func (UnionString) ImplementsGlobalSpendGetReportResponseMessagesUnion()                           {}
func (UnionString) ImplementsGlobalSpendGetReportResponseResponseUnion()                           {}
func (UnionString) ImplementsGlobalSpendGetReportResponseStartTimeUnion()                          {}
func (UnionString) ImplementsGuardrailListResponseGuardrailsLlmParamsModeUnion()                   {}

type UnionInt int64

func (UnionInt) ImplementsFineTuningJobNewParamsHyperparametersBatchSizeUnion() {}
func (UnionInt) ImplementsFineTuningJobNewParamsHyperparametersNEpochsUnion()   {}

type UnionFloat float64

func (UnionFloat) ImplementsModelNewParamsLlmParamsStreamTimeoutUnion()                        {}
func (UnionFloat) ImplementsModelNewParamsLlmParamsTimeoutUnion()                              {}
func (UnionFloat) ImplementsUpdateDeploymentLlmParamsStreamTimeoutUnionParam()                 {}
func (UnionFloat) ImplementsUpdateDeploymentLlmParamsTimeoutUnionParam()                       {}
func (UnionFloat) ImplementsFineTuningJobNewParamsHyperparametersLearningRateMultiplierUnion() {}

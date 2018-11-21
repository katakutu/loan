package helper

import (
	"github.com/astaxie/beego"

	"micro-loan/common/i18n"
	"micro-loan/common/pkg/admin"
	"micro-loan/common/pkg/feedback"
	"micro-loan/common/pkg/rbac"
	"micro-loan/common/pkg/ticket"
	"micro-loan/common/service"
	"micro-loan/common/tools"
)

func init() {
	beego.AddFuncMap("getDateMHS", tools.GetDateMHS)
	beego.AddFuncMap("dateMHSZ", tools.DateMHSZ)
	beego.AddFuncMap("mDateMH", tools.MDateMH)
	beego.AddFuncMap("mDateMHS", tools.MDateMHS)
	beego.AddFuncMap("mDateMHSDate", tools.MDateMHSDate)
	beego.AddFuncMap("mDateUTC", tools.MDateUTC)
	beego.AddFuncMap("calcDateRange", tools.GetDateRangeMillis)

	beego.AddFuncMap("statusDisplay", service.StatusDisplay)
	beego.AddFuncMap("typeDisplayProduct", service.TypeDisplayProduct)
	beego.AddFuncMap("statusDisplayProduct", service.StatusDisplayProduct)
	beego.AddFuncMap("statusDisplayProductRepay", service.StatusDisplayProductRepay)
	beego.AddFuncMap("getCustomerTags", service.GetCustomerTags)
	beego.AddFuncMap("getResourceUserMark", service.GetResourceUserMark)
	beego.AddFuncMap("getLoanStatusDesc", service.GetLoanStatusDesc)
	beego.AddFuncMap("getPayTypeDesc", service.GetPayTypeDesc)
	beego.AddFuncMap("getVaCompanyTypeDesc", service.GetVaCompanyTypeDesc)
	beego.AddFuncMap("getGenderDisplay", service.GetGenderDisplay)
	beego.AddFuncMap("genImgHTML", service.GenImgHTML)
	beego.AddFuncMap("genImgHref", service.GenImgHref)
	beego.AddFuncMap("tplTimeNow", service.TplTimeNow)
	beego.AddFuncMap("operatorName", admin.OperatorName)
	beego.AddFuncMap("riskTypeDisplay", service.RiskTypeDisplay)
	beego.AddFuncMap("riskItemDisplay", service.RiskItemDisplay)
	beego.AddFuncMap("riskReasonDisplay", service.RiskReasonDisplay)
	beego.AddFuncMap("riskRelieveReasonDisplay", service.RiskRelieveReasonDisplay)
	beego.AddFuncMap("riskCtlStatusDisplay", service.RiskCtlStatusDisplay)
	beego.AddFuncMap("riskCtlOperateCmd", service.RiskCtlOperateCmd)
	beego.AddFuncMap("getIsLoan", service.GetIsReloan)
	beego.AddFuncMap("getOpLoggerCodeDesc", service.GetOpLoggerCodeDesc)
	beego.AddFuncMap("serviceTypeDisplay", service.VerfiyCodeServiceTypeDisplay)
	beego.AddFuncMap("authCodeTypeDisplay", service.AuthCodeTypeDisplay)
	beego.AddFuncMap("expireTimeCalculate", service.ExpireTimeCalculate)
	beego.AddFuncMap("urgeOutReasonDisplay", service.UrgeOutReasonDisplay)
	beego.AddFuncMap("systemConfigItemTypeDisplay", service.SystemConfigItemTypeDisplay)
	beego.AddFuncMap("smsVerifyCodeStatusDisplay", service.SmsVerifyCodeStatusDisplay)
	beego.AddFuncMap("phoneConnectDisplay", service.PhoneConnectDisplay)
	beego.AddFuncMap("repayInclinationDisplay", service.RepayInclinationDisplay)
	beego.AddFuncMap("unconnectReasonDisplay", service.UnconnectReasonDisplay)
	beego.AddFuncMap("overdueReasonItemDisplay", service.OverdueReasonItemDisplay)
	beego.AddFuncMap("riskStatusDisplay", service.RiskStatusDisplay)
	beego.AddFuncMap("smsServiceDisplay", service.SmsServiceDisplay)
	beego.AddFuncMap("smsServiceTypeDisplay", service.SmsServiceTypeDisplay)
	beego.AddFuncMap("smsDeliveryStatusDisplay", service.SmsDeliveryStatusDisplay)
	beego.AddFuncMap("phoneObjectDisplay", service.PhoneObjectDisplay)
	beego.AddFuncMap("urgeCallTypeDisplay", service.UrgeCallTypeDisplay)
	beego.AddFuncMap("isInMap", service.IsInMap)
	beego.AddFuncMap("getFeedbackTagDisplay", feedback.GetTagDisplay)
	beego.AddFuncMap("getThirdpartyName", service.GetThirdpartyNameForTemplate)
	beego.AddFuncMap("displayLimitText", service.DisplayLimitText)
	beego.AddFuncMap("getRoleNameDisplay", service.GetRoleNameDisplay)
	beego.AddFuncMap("getRoleTypeDisplay", service.GetRoleTypeDisplay)
	beego.AddFuncMap("getRoleName", rbac.GetRoleName)
	beego.AddFuncMap("getTicketItemDisplay", service.GetTicketItemDisplay)
	beego.AddFuncMap("getTicketStatusDisplay", service.GetTicketStatusDisplay)
	beego.AddFuncMap("getCommunicationWayDisplay", service.GetCommunicationWayDisplay)
	beego.AddFuncMap("getIsEmptyDisplay", service.GetIsEmptyDisplay)
	beego.AddFuncMap("displayWorkerCanAcceptTicket", ticket.DisplayWorkerCanAcceptTicket)
	beego.AddFuncMap("getRiskLevelDisplay", ticket.GetRiskLevelDisplay)
	beego.AddFuncMap("isTicketFinished", ticket.IsTicketFinished)
	beego.AddFuncMap("chargeFeeTypeDisplay", service.ChargeFeeTypeDisplay)
	beego.AddFuncMap("ceilWayDisplay", service.CeilWayDisplay)
	beego.AddFuncMap("productOptTypeDisplay", service.ProductOptTypeDisplay)
	beego.AddFuncMap("chargeTypeDisplay", service.ChargeTypeDisplay)
	beego.AddFuncMap("getThirdpartyName", service.GetThirdpartyName)
	beego.AddFuncMap("overdueDaysDisplay", service.OverdueDaysDisplay)
	beego.AddFuncMap("isOutDisplay", service.IsOutDisplay)
	beego.AddFuncMap("urgeTypeDisplay", service.UrgeTypeDisplay)
	beego.AddFuncMap("entrustStatusDisplay", service.EntrustStatusDisplay)
	beego.AddFuncMap("entrustCompanyDisplay", service.EntrustCompanyDisplay)
	beego.AddFuncMap("paymentCodeDisplay", service.PaymentCodeDisplay)
	beego.AddFuncMap("marketPaymentCodeGenerateButton", service.MarketPaymentCodeGenerateButton)
	beego.AddFuncMap("platformMarkDisplay", service.PlatformMarkDisplay)
	beego.AddFuncMap("checkPromiseIsToday", service.CheckPromiseIsToday)
	beego.AddFuncMap("reduceTypeDisplay", service.ReduceTypeDisplay)
	beego.AddFuncMap("reduceStatusDisplay", service.ReduceStatusDisplay)
	beego.AddFuncMap("voipCallMothodDisplay", service.VoipCallMothodDisplay)
	beego.AddFuncMap("voipCallDirectionDisplay", service.VoipCallDirectionDisplay)
	beego.AddFuncMap("voipIsDialDisplay", service.VoipIsDialDisplay)
	beego.AddFuncMap("voipHangupDisplay", service.VoipHangupDisplay)
	beego.AddFuncMap("couponTypeDisplay", service.CouponTypeDisplay)
	beego.AddFuncMap("couponAvaliableDisplay", service.CouponAvaliableDisplay)
	beego.AddFuncMap("couponDistributeDisplay", service.CouponDistributeDisplay)
	beego.AddFuncMap("couponStatusDisplay", service.CouponStatusDisplay)
	beego.AddFuncMap("phoneVerifyCallResult", service.PhoneVerifyCallResult)
	beego.AddFuncMap("containNumber", tools.ContainNumber)
	beego.AddFuncMap("getFailedDisburseOrderReason", service.GetFailedDisburseOrderReason)
	beego.AddFuncMap("vaDisplayBankCode", service.VaDisplayBankCode)
	beego.AddFuncMap("companyDisplay", service.CompanyDisplay)
	beego.AddFuncMap("repaymentSource", service.RepaymentSourceDisplay)
	beego.AddFuncMap("companyDisplayByCode", service.CompanyDisplayByCode)

	beego.AddFuncMap("pushTargetDisplay", service.PushTargetDisplay)
	beego.AddFuncMap("messageTypeDisplay", service.MessageTypeDisplay)
	beego.AddFuncMap("pushWayDisplay", service.PushWayDisplay)
	beego.AddFuncMap("schemaModeDisplay", service.SchemaModeDisplay)
	beego.AddFuncMap("schemaStatusDisplay", service.SchemaStatusDisplay)
	beego.AddFuncMap("couponTargetDisplay", service.CouponTargetDisplay)
	beego.AddFuncMap("adPositionDisplay", service.AdPositionDisplay)
	beego.AddFuncMap("smsTargetDisplay", service.SmsTargetDisplay)

	beego.AddFuncMap("bannerTypeDisplay", service.BannerTypeDisplay)

	beego.AddFuncMap("numberFormat", tools.NumberFormat)
	beego.AddFuncMap("t", i18n.T)
}

package routers

import (
	"github.com/astaxie/beego"

	"micro-loan/admin/app/controllers"
	"micro-loan/common/cprof"
)

func init() {
	beego.Router("/health_checker", &controllers.MainController{}, "*:HealthChecker")

	beego.Router("/", &controllers.MainController{})
	beego.Router("/debug", &controllers.DebugController{})
	beego.Router("/ping", &controllers.MainController{}, "*:Ping")
	beego.Router("/crypto", &controllers.MainController{}, "*:Crypto")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/login_confirm", &controllers.MainController{}, "*:LoginConfirm")

	beego.Router("/index", &controllers.IndexController{}, "*:Dashboard")
	beego.Router("/sms/sms_status_list", &controllers.SmsController{}, "get:SmsStatusList")

	beego.Router("/product/list", &controllers.ProductController{}, "*:List")
	beego.Router("/product/edit", &controllers.ProductController{}, "*:Edit")
	beego.Router("/product/add", &controllers.ProductController{}, "*:Add")
	beego.Router("/product/doedit", &controllers.ProductController{}, "*:DoEdit")
	beego.Router("/product/trial_calc", &controllers.ProductController{}, "*:TrialCalc")
	beego.Router("/product/down", &controllers.ProductController{}, "*:Down")
	beego.Router("/product/up", &controllers.ProductController{}, "*:Up")
	beego.Router("/product/clone", &controllers.ProductController{}, "*:Clone")
	beego.Router("/product/opt/record", &controllers.ProductController{}, "*:OptRecord")
	beego.Router("/product/opt/record/view", &controllers.ProductController{}, "get:OptRecordView")

	beego.Router("/thirdparty/add", &controllers.ThirdPartyController{}, "*:Add")
	beego.Router("/thirdparty/list", &controllers.ThirdPartyController{}, "*:List")
	beego.Router("/thirdparty/edit", &controllers.ThirdPartyController{}, "*:Edit")
	beego.Router("/thirdparty/doedit", &controllers.ThirdPartyController{}, "*:DoEdit")
	beego.Router("/thirdparty/reconciliation", &controllers.ThirdPartyController{}, "*:Reconciliation")
	beego.Router("/thirdparty/reconciliation_customer", &controllers.ThirdPartyController{}, "*:ReconciliationCustomer")
	beego.Router("/thirdparty/reconciliation_customer_detail", &controllers.ThirdPartyController{}, "get:ReconciliationCustomerDetail")

	beego.Router("/manage/admin/list", &controllers.ManageController{}, "*:AdminList")
	beego.Router("/manage/admin/create", &controllers.ManageController{}, "*:AdminCreate")
	beego.Router("/manage/admin/save", &controllers.ManageController{}, "*:AdminSave")
	beego.Router("/manage/admin/block", &controllers.ManageController{}, "*:AdminBlock")
	beego.Router("/manage/admin/unblock", &controllers.ManageController{}, "*:AdminUnblock")
	beego.Router("/manage/admin/edit", &controllers.ManageController{}, "get:AdminEdit")
	beego.Router("/manage/admin/update", &controllers.ManageController{}, "post:AdminUpdate")
	beego.Router("/manage/sms_verify_code", &controllers.ManageController{}, "get:SmsVerifyCode")
	beego.Router("/manage/order/change", &controllers.ManageController{}, "get:OrderChange")
	beego.Router("/manage/order/change/save", &controllers.ManageController{}, "post:OrderChangeSave")

	beego.Router("/sysconfig/system_config/list", &controllers.SysConfigController{}, "get:SystemConfigList")
	beego.Router("/sysconfig/system_config/save", &controllers.SysConfigController{}, "post:SystemConfigSave")

	beego.Router("/menu", &controllers.MenuController{}, "get:List")
	beego.Router("/menu/create", &controllers.MenuController{}, "post:Create")
	beego.Router("/menu/update_page", &controllers.MenuController{}, "get:UpdatePage")
	beego.Router("/menu/update", &controllers.MenuController{}, "post:Update")
	beego.Router("/menu/delete", &controllers.MenuController{}, "post:Delete")
	beego.Router("/menu/update_sort", &controllers.MenuController{}, "post:UpdateSort")

	beego.Router("/rbac/role", &controllers.RBACController{}, "get:RoleList")
	beego.Router("/rbac/role/create", &controllers.RBACController{}, "post:RoleCreate")
	beego.Router("/rbac/role/update_page", &controllers.RBACController{}, "get:RoleEditPage")
	beego.Router("/rbac/role/update", &controllers.RBACController{}, "post:RoleUpdate")
	beego.Router("/rbac/role/privilege/manage", &controllers.RBACController{}, "get:RolePrivilegeManage")
	beego.Router("/rbac/role/privileges", &controllers.RBACController{}, "get:RolePrivileges")
	beego.Router("/rbac/role/assign_privileges", &controllers.RBACController{}, "post:RoleAssignPrivileges")
	beego.Router("/rbac/role/revoke_privileges", &controllers.RBACController{}, "post:RoleRevokePrivileges")

	beego.Router("/rbac/operation", &controllers.RBACController{}, "get:OperationList")
	beego.Router("/rbac/operation/create", &controllers.RBACController{}, "post:OperationCreate")
	beego.Router("/rbac/operation/update_page", &controllers.RBACController{}, "get:OperationUpdatePage")
	beego.Router("/rbac/operation/update", &controllers.RBACController{}, "post:OperationUpdate")

	beego.Router("/fund/manage/loan", &controllers.FundController{}, "get:LoanConfig")
	beego.Router("/fund/manage/repay", &controllers.FundController{}, "get:RepayConfig")
	beego.Router("/fund/manage/bank/query", &controllers.FundController{}, "*:BankQuery")
	beego.Router("/fund/manage/bank/assign", &controllers.FundController{}, "*:BankAssign")
	beego.Router("/fund/manage/bank/un_assign", &controllers.FundController{}, "*:BankUnAssign")

	beego.Router("/rbac/privilege", &controllers.RBACController{}, "get:PrivilegeList")
	beego.Router("/rbac/privilege/create", &controllers.RBACController{}, "post:PrivilegeCreate")
	beego.Router("/rbac/privilege/update_page", &controllers.RBACController{}, "get:PrivilegeUpdatePage")
	beego.Router("/rbac/privilege/update", &controllers.RBACController{}, "post:PrivilegeUpdate")
	beego.Router("/rbac/privilege/operation/manage", &controllers.RBACController{}, "get:PrivilegeOperationManage")
	beego.Router("/rbac/privilege/operations", &controllers.RBACController{}, "get:PrivilegeOperations")
	beego.Router("/rbac/privilege/assign_operations", &controllers.RBACController{}, "post:PrivilegeAssignOperations")
	beego.Router("/rbac/privilege/revoke_operations", &controllers.RBACController{}, "post:PrivilegeRevokeOperations")

	beego.Router("/rbac/privilege_group", &controllers.RBACController{}, "get:PrivilegeGroupList")
	beego.Router("/rbac/privilege_group/create", &controllers.RBACController{}, "post:PrivilegeGroupCreate")
	beego.Router("/rbac/privilege_group/update_page", &controllers.RBACController{}, "get:PrivilegeGroupUpdatePage")
	beego.Router("/rbac/privilege_group/update", &controllers.RBACController{}, "post:PrivilegeGroupUpdate")

	beego.Router("/ticket/list", &controllers.TicketController{}, "get:ManageList")
	beego.Router("/ticket/me", &controllers.TicketController{}, "get:Me")
	beego.Router("/ticket/collection", &controllers.TicketController{}, "get:Collection")
	beego.Router("/ticket/rm", &controllers.TicketController{}, "get:RmTicket")
	beego.Router("/ticket/pv_inforeview", &controllers.TicketController{}, "get:PvAndInfoReview")
	beego.Router("/ticket/update_status", &controllers.TicketController{}, "post:UpdateStatus")
	beego.Router("/ticket/assign_page", &controllers.TicketController{}, "get:AssignPage")
	beego.Router("/ticket/assign", &controllers.TicketController{}, "post:Assign")
	beego.Router("/ticket/batch_assign_page", &controllers.TicketController{}, "get:BatchAssignPage")
	beego.Router("/ticket/batch_assign", &controllers.TicketController{}, "post:BatchAssign")
	beego.Router("/ticket/batch_applyentrust", &controllers.TicketController{}, "post:BatchApplyEntrust")
	beego.Router("/ticket/worker_manage", &controllers.TicketController{}, "get:WorkerManage")
	beego.Router("/ticket/ajax_modify_reduced_quota", &controllers.TicketController{}, "post:AjaxModifyReducedQuota")
	beego.Router("/ticket/update_worker_status", &controllers.TicketController{}, "post:UpdateWorkerStatus")
	beego.Router("/ticket/update_worker_online_status", &controllers.TicketController{}, "post:UpdateWorkerOnlineStatus")
	beego.Router("/ticket/update_my_online_status", &controllers.TicketController{}, "post:UpdateMyOnlineStatus")
	beego.Router("/ticket/performance/me", &controllers.TicketController{}, "get:PerformanceMe")
	beego.Router("/ticket/performance/management", &controllers.TicketController{}, "get:PerformanceManagement")
	beego.Router("/ticket/performance/daily/export", &controllers.TicketController{}, "get:DailyPerformanceExport")
	beego.Router("/ticket/performance/management_hour", &controllers.TicketController{}, "get:PerformanceManagementHour")
	beego.Router("/ticket/item/performance/month", &controllers.TicketController{}, "get:ItemPerformanceMonthStats")
	beego.Router("/ticket/process/me", &controllers.TicketController{}, "get:MyProcess")

	beego.Router("/business/detail/list", &controllers.BusinessDetailController{}, "*:List")
	beego.Router("/business/detail/recharge", &controllers.BusinessDetailController{}, "*:Recharge")
	beego.Router("/business/detail/recharge/save", &controllers.BusinessDetailController{}, "*:DoSaveRecharge")
	beego.Router("/business/detail/withdraw", &controllers.BusinessDetailController{}, "*:Withdraw")
	beego.Router("/business/detail/withdraw/save", &controllers.BusinessDetailController{}, "*:DoWithdraw")
	beego.Router("/business/detail/list/view", &controllers.BusinessDetailController{}, "*:ListDetail")

	beego.Router("/customer/list", &controllers.CustomerController{}, "*:List")
	beego.Router("/customer/refund/apply", &controllers.CustomerController{}, "*:RefundApply")
	beego.Router("/customer/refund/do", &controllers.CustomerController{}, "post:DoRefund")
	beego.Router("/customer/modify/mobile", &controllers.CustomerController{}, "*:ModifyMobile")
	beego.Router("/customer/modify/mobile/do", &controllers.CustomerController{}, "*:DoModifyMobile")
	beego.Router("/customer/pic_show", &controllers.CustomerController{}, "*:PicShow")
	beego.Router("/customer/risk", &controllers.CustomerController{}, "*:Risk")
	beego.Router("/customer/detail", &controllers.CustomerController{}, "*:Detail")
	// 客户资料相关权限细化 {
	beego.Router("/customer/detail/base-info", &controllers.CustomerController{}, "*:DetailBaseInfo")
	beego.Router("/customer/detail/other-info", &controllers.CustomerController{}, "*:DetailOtherInfo")
	beego.Router("/customer/detail/big-data-info", &controllers.CustomerController{}, "*:DetailBigDataInfo")
	beego.Router("/customer/detail/communication-record", &controllers.CustomerController{}, "*:DetailCommunicationRecord")
	beego.Router("/customer/detail/loan-history", &controllers.CustomerController{}, "*:DetailLoanHistory")
	beego.Router("/customer/detail/check-duplicate", &controllers.CustomerController{}, "*:DetailCheckDuplicate")
	// }
	beego.Router("/customer/follow", &controllers.CustomerController{}, "*:Follow")
	beego.Router("/customer/follow/confirm", &controllers.CustomerController{}, "post:FollowConfirm")
	beego.Router("/customer/risk_report", &controllers.CustomerController{}, "*:RiskReport")
	beego.Router("/customer/risk_query_val", &controllers.CustomerController{}, "get:RiskQueryVal")
	beego.Router("/customer/risk_relieve", &controllers.CustomerController{}, "*:RiskRelieve")
	beego.Router("/customer/risk_report/save", &controllers.CustomerController{}, "post:RiskSave")
	beego.Router("/customer/ajax_modify", &controllers.CustomerController{}, "*:AjaxModify")
	beego.Router("/customer/risk_relieve/save", &controllers.CustomerController{}, "post:RiskRelieveSave")
	beego.Router("/customer/risk_review", &controllers.CustomerController{}, "*:RiskReview")
	beego.Router("/customer/risk_review/save", &controllers.CustomerController{}, "post:RiskReviewSave")
	beego.Router("/customer/delete", &controllers.CustomerController{}, "*:Delete")
	beego.Router("/customer/super_delete", &controllers.CustomerController{}, "*:SuperDelete")
	beego.Router("/customer/import_blacklist", &controllers.CustomerController{}, "*:ImportBlacklist")
	beego.Router("/customer/blacklist_save", &controllers.CustomerController{}, "post:BlacklistSave")

	beego.Router("/riskctl/list", &controllers.RiskCtlController{}, "*:List")
	beego.Router("/riskctl/regular/all", &controllers.RiskCtlController{}, "post:RegularAll")
	beego.Router("/riskctl/phone_verify", &controllers.RiskCtlController{}, "*:PhoneVerify")
	beego.Router("/riskctl/phone_verify/save", &controllers.RiskCtlController{}, "post:PhoneVerifySave")
	beego.Router("/riskctl/phone_verify/call_record", &controllers.RiskCtlController{}, "*:PhoneVerifyCallRecord")
	beego.Router("/riskctl/phone_verify/detail", &controllers.RiskCtlController{}, "*:PhoneVerifyCallDetail")
	beego.Router("/riskctl/show_verify_result", &controllers.RiskCtlController{}, "post:ShowVerifyResult")
	beego.Router("/riskctl/check_blacklist", &controllers.RiskCtlController{}, "get:CheckBlacklist")
	beego.Router("/riskctl/follow", &controllers.RiskCtlController{}, "*:Follow")

	//这一组路由都为大数据提供服务了
	beego.Router("/risknotify/notify", &controllers.RiskNotifyController{}, "post:Notify")
	beego.Router("/risknotify/quota_conf", &controllers.RiskNotifyController{}, "post:QuotaConf")
	beego.Router("/risknotify/thirdparty_query", &controllers.RiskNotifyController{}, "post:ThirdPartyQuery")
	beego.Router("/risknotify/risk_query", &controllers.RiskNotifyController{}, "post:RiskQuery")

	beego.Router("/order/list", &controllers.OrderController{}, "*:List")
	beego.Router("/order/backend/business_history", &controllers.OrderController{}, "*:BusinessHistory")
	beego.Router("/order/backend/repay_plan", &controllers.OrderController{}, "*:RepayPlan")

	beego.Router("/loan/list", &controllers.LoanController{}, "*:List")
	beego.Router("/loan/backend/edit_bank_info", &controllers.LoanController{}, "*:EditBankInfo")
	beego.Router("/loan/backend/do_edit_bank_info", &controllers.LoanController{}, "*:DoEditBankInfo")
	beego.Router("/loan/backend/disburse_again", &controllers.LoanController{}, "*:DisbureAgain")
	beego.Router("/loan/backend/do_disburse_again", &controllers.LoanController{}, "*:DoDisbureAgain")
	beego.Router("/loan/backend/do/disbure/again/multi", &controllers.LoanController{}, "*:DoDisbureAgainMulti")
	beego.Router("/loan/backend/do/roll/back", &controllers.LoanController{}, "*:DoRollBack")

	beego.Router("/repay/list", &controllers.RepayController{}, "*:List")
	beego.Router("/repay/va/search", &controllers.RepayController{}, "*:VaSearch")
	beego.Router("/repay/backend/repay_plan", &controllers.RepayController{}, "*:RepayPlan")
	beego.Router("/repay/backend/repay_plan/rollback", &controllers.RepayController{}, "*:RepayPlanRollBack")
	beego.Router("/repay/backend/repay_plan/rollback/do", &controllers.RepayController{}, "*:DoRepayPlanRollBack")
	beego.Router("/repay/backend/user_e_trans", &controllers.RepayController{}, "*:UserTrans")
	beego.Router("/repay/backend/repay_plan_history", &controllers.RepayController{}, "*:RepayPlanHistory")
	beego.Router("/repay/remind_case/list", &controllers.RepayController{}, "*:RemindCaseList")
	beego.Router("/repay/remind_case/handle", &controllers.RepayController{}, "*:RemindCaseHandle")
	beego.Router("/repay/remind_case/update", &controllers.RepayController{}, "*:RemindCaseUpdate")
	beego.Router("/repay/remind_case/view", &controllers.RepayController{}, "*:RemindCaseView")
	beego.Router("/repay/remind_case/log", &controllers.RepayController{}, "*:RemindCaseLog")
	beego.Router("/overdue/list", &controllers.OverdueController{}, "*:List")
	beego.Router("/overdue/urge", &controllers.OverdueController{}, "*:Urge")
	beego.Router("/overdue/urge/save", &controllers.OverdueController{}, "post:UrgeSave")
	beego.Router("/overdue/urge/prereduced", &controllers.OverdueController{}, "post:PreReduced")
	beego.Router("/overdue/urge/apply_entrust", &controllers.OverdueController{}, "post:ApplyEntrust")
	beego.Router("/overdue/urge/detail", &controllers.OverdueController{}, "*:UrgeDetail")
	beego.Router("/overdue/backend/reduction/list", &controllers.OverdueController{}, "*:ReductionList")
	beego.Router("/overdue/backend/reduction/confirm", &controllers.OverdueController{}, "*:ReductionConfirm")
	beego.Router("/overdue/backend/reduction/confirm/save", &controllers.OverdueController{}, "*:ReductionConfirmSave")
	beego.Router("/overdue/backend/reduction", &controllers.OverdueController{}, "*:Reduction")
	beego.Router("/overdue/backend/do_reduction", &controllers.OverdueController{}, "*:DoReduction")
	beego.Router("/overdue/defer_show", &controllers.OverdueController{}, "*:DeferShow")
	beego.Router("/overdue/roll_trial_data", &controllers.OverdueController{}, "get:GetRollTrialData")
	beego.Router("/overdue/market_payment_code/generate", &controllers.OverdueController{}, "*:MarketPaymentCodeGenerate")
	beego.Router("/overdue/roll_apply", &controllers.OverdueController{}, "*:RollApply")
	beego.Router("/overdue/co2case_list", &controllers.OverdueController{}, "*:Co2caseList")
	beego.Router("/overdue/entrust_approval_list", &controllers.OverdueController{}, "*:EntrustApprovalList")
	beego.Router("/overdue/entrust_approval/batch_assign_page", &controllers.OverdueController{}, "get:BatchAssignPage")
	beego.Router("/overdue/entrust_approval/batch_assign", &controllers.OverdueController{}, "post:BatchAssign")
	beego.Router("/overdue/create_ticket", &controllers.OverdueController{}, "*:CreateTicket")
	// beego.Router("/overdue/entrust_apply", &controllers.OverdueController{}, "*:EntrustApply")

	beego.Router("/autocall/record", &controllers.AutoCallController{}, "*:AutoCallRecord")

	beego.Router("/feedback", &controllers.FeedbackController{}, "*:List")
	beego.Router("/feedback/export", &controllers.FeedbackController{}, "*:Export")
	beego.Router("/feedback/image", &controllers.FeedbackController{}, "*:Image")

	beego.Router("/feedback/paymentvocher", &controllers.FeedbackController{}, "*:PaymentVocherList")
	beego.Router("/feedback/payment_vocher_image", &controllers.FeedbackController{}, "*:PaymentVocherImage")
	beego.Router("/feedback/update_payment_vocher", &controllers.FeedbackController{}, "*:UpdatePaymentVocher")

	beego.Router("/account/password", &controllers.AdminController{}, "*:Password")
	beego.Router("/account/fixpassword", &controllers.AdminController{}, "post:FixPassword")
	beego.Router("/admin/op_log", &controllers.AdminController{}, "get:OpLog")
	beego.Router("/admin/op_log/view", &controllers.AdminController{}, "get:OpLogView")
	beego.Router("/admin/list", &controllers.AdminController{}, "get:List")
	beego.Router("/admin/edit", &controllers.AdminController{}, "get:Edit")
	beego.Router("/admin/create", &controllers.AdminController{}, "get:Create")
	beego.Router("/admin/save", &controllers.AdminController{}, "post:Save")
	beego.Router("/admin/update", &controllers.AdminController{}, "*:Update")
	beego.Router("/admin/block", &controllers.AdminController{}, "*:Block")
	beego.Router("/admin/unblock", &controllers.AdminController{}, "*:Unblock")
	beego.Router("/admin/export/:id", &controllers.AdminController{}, "*:Export")

	beego.Router("/resource/img/:rid", &controllers.ResourceController{}, "get:FetchImgStream")
	beego.Router("/resource/media/:rid", &controllers.ResourceController{}, "get:FetchAudioStream")

	beego.Router("/monitor", &controllers.MonitorController{}, "*:Monitor")
	beego.Router("/monitor/list", &controllers.MonitorController{}, "*:List")

	beego.Router("/thirdparty_record/list", &controllers.ThirdPartyRecordController{}, "*:List")
	beego.Router("/thirdparty_record/detail", &controllers.ThirdPartyRecordController{}, "*:Detail")

	beego.Router("/transaction_inquiry/disburse_inquiry_form", &controllers.TransactionInquiryController{}, "*:DisburseInquiryForm")
	beego.Router("/transaction_inquiry/disburse_inquiry_result", &controllers.TransactionInquiryController{}, "*:DisburseInquiryResult")
	beego.Router("/transaction_inquiry/payment_inquiry_form", &controllers.TransactionInquiryController{}, "*:PaymentInquiryForm")
	beego.Router("/transaction_inquiry/payment_inquiry_result", &controllers.TransactionInquiryController{}, "*:PaymentInquiryResult")
	beego.Router("/transaction_inquiry/generate_fix_doku", &controllers.TransactionInquiryController{}, "*:GenerateFixDoku")

	beego.Router("/extension/sip_call", &controllers.ExtensionController{}, "*:SipCall")
	beego.Router("/extension/sip_call_result", &controllers.ExtensionController{}, "*:SipCallResult")
	beego.Router("/extension/sip_call_bill", &controllers.ExtensionController{}, "*:SipCallBill")
	beego.Router("/extension/list", &controllers.ExtensionController{}, "*:List")
	beego.Router("/extension/ext_history", &controllers.ExtensionController{}, "*:ExtHistory")
	beego.Router("/extension/call_record", &controllers.ExtensionController{}, "*:CallRecord")
	beego.Router("/extension/update_ext_info", &controllers.ExtensionController{}, "*:UpdateExtInfo")
	beego.Router("/extension/assign_page", &controllers.ExtensionController{}, "get:AssignPage")
	beego.Router("/extension/assign", &controllers.ExtensionController{}, "post:Assign")
	beego.Router("/extension/cancelassign", &controllers.ExtensionController{}, "*:CancelAssign")

	beego.Router("/coupon", &controllers.CouponController{}, "*:Coupon")
	beego.Router("/coupon/list", &controllers.CouponController{}, "*:CouponList")
	beego.Router("/coupon/edit", &controllers.CouponController{}, "*:CouponEdit")
	beego.Router("/coupon/edit_save", &controllers.CouponController{}, "*:CouponEditSave")
	beego.Router("/coupon/active", &controllers.CouponController{}, "*:CouponActive")
	beego.Router("/coupon/list_export", &controllers.CouponController{}, "*:CouponListExport")
	beego.Router("/coupon/coupon_detail", &controllers.CouponController{}, "*:CouponDetail")

	beego.Router("/operation/list_advertisement", &controllers.OperationController{}, "*:ListAdvertisement")
	beego.Router("/operation/add_advertisement", &controllers.OperationController{}, "*:AddAdvertisement")
	beego.Router("/operation/update_advertisement", &controllers.OperationController{}, "*:UpdateAdvertisement")
	beego.Router("/operation/del_advertisement", &controllers.OperationController{}, "*:DelAdvertisement")

	beego.Router("/operation/list_banner", &controllers.OperationController{}, "*:ListBanner")
	beego.Router("/operation/add_banner", &controllers.OperationController{}, "*:AddBanner")
	beego.Router("/operation/del_banner", &controllers.OperationController{}, "*:DelBanner")
	beego.Router("/operation/update_banner_postion", &controllers.OperationController{}, "*:UpdateBannerPostion")

	beego.Router("/operation/list_ad_position", &controllers.OperationController{}, "*:ListAdPosition")
	beego.Router("/operation/add_ad_position", &controllers.OperationController{}, "*:AddAdPosition")
	beego.Router("/operation/del_ad_position", &controllers.OperationController{}, "*:DelAdPosition")

	// pprof
	beego.Router("/debug/pprof", &cprof.ProfController{}, "*:Get")
	beego.Router(`/debug/pprof/:pp([\w]+)`, &cprof.ProfController{}, "*:Get")

	beego.Router("/schema/run", &controllers.SchemaController{}, "*:SchemaRun")

	beego.Router("/schema/push_list", &controllers.SchemaController{}, "*:PushList")
	beego.Router("/schema/push_edit", &controllers.SchemaController{}, "*:PushEdit")
	beego.Router("/schema/push_save", &controllers.SchemaController{}, "*:PushSave")
	beego.Router("/schema/push_detail", &controllers.SchemaController{}, "*:PushDetail")
	beego.Router("/schema/push_active", &controllers.SchemaController{}, "*:PushActive")

	beego.Router("/schema/sms_list", &controllers.SchemaController{}, "*:SmsList")
	beego.Router("/schema/sms_edit", &controllers.SchemaController{}, "*:SmsEdit")
	beego.Router("/schema/sms_save", &controllers.SchemaController{}, "*:SmsSave")
	beego.Router("/schema/sms_detail", &controllers.SchemaController{}, "*:SmsDetail")
	beego.Router("/schema/sms_active", &controllers.SchemaController{}, "*:SmsActive")

	beego.Router("/schema/coupon_list", &controllers.SchemaController{}, "*:CouponList")
	beego.Router("/schema/coupon_edit", &controllers.SchemaController{}, "*:CouponEdit")
	beego.Router("/schema/coupon_save", &controllers.SchemaController{}, "*:CouponSave")
	beego.Router("/schema/coupon_active", &controllers.SchemaController{}, "*:CouponActive")

	beego.Router("/activity/list_floating", &controllers.ActivityController{}, "*:ListFloating")
	beego.Router("/activity/add_floating", &controllers.ActivityController{}, "*:AddFloating")
	beego.Router("/activity/del_floating", &controllers.ActivityController{}, "*:DelFloating")

	beego.Router("/activity/list_popupwindow", &controllers.ActivityController{}, "*:ListPopUpWindow")
	beego.Router("/activity/add_popupwindow", &controllers.ActivityController{}, "*:AddPopUpWindow")
	beego.Router("/activity/del_popupwindow", &controllers.ActivityController{}, "*:DelPopUpWindow")
}

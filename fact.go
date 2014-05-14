type Fact struct{
  fact_transaction_id int64
  rule_id int64
  rule_version_number int64
  customer_id int64
  merchandiser_id int64
  merchandiser_manager_id int64
  order_id int64
  legacy_order_number string
  type string
}
/*
    t.integer  "team_id"
    t.integer  "period_id"
    t.decimal  "amount",                               :precision => 10, :scale => 2
    t.string   "sku"
    t.string   "code"
    t.datetime "date"
    t.datetime "created_at",                                                           :null => false
    t.string   "authorize_id"
    t.string   "rma_number"
    t.text     "notes"
    t.string   "sales_channel"
    t.integer  "caused_by"
    t.integer  "applied_to"
    t.integer  "credit_id"
    t.integer  "user_id"
    t.integer  "popup_id"
    t.integer  "rma_id"
    t.boolean  "fast_start"
    t.string   "threshold_target1"
    t.decimal  "threshold_number1",                    :precision => 10, :scale => 2
    t.string   "threshold_string1"
    t.string   "threshold_period1"
    t.string   "threshold_target2"
    t.decimal  "threshold_number2",                    :precision => 10, :scale => 2
    t.string   "threshold_string2"
    t.string   "threshold_period2"
    t.string   "threshold_target3"
    t.decimal  "threshold_number3",                    :precision => 10, :scale => 2
    t.string   "threshold_string3"
    t.string   "threshold_period3"
    t.string   "threshold_target4"
    t.decimal  "threshold_number4",                    :precision => 10, :scale => 2
    t.string   "threshold_string4"
    t.string   "threshold_period4"
    t.string   "threshold_target5"
    t.decimal  "threshold_number5",                    :precision => 10, :scale => 2
    t.string   "threshold_string5"
    t.string   "threshold_period5"
    t.datetime "updated_at",                                                           :null => false
    t.integer  "billing_address_id"
    t.string   "color"
    t.integer  "credit_card_id"
    t.datetime "effective_date"
    t.datetime "expired_at"
    t.integer  "hostess_id"
    t.string   "ip_address"
    t.string   "item_type"
    t.string   "label"
    t.decimal  "percent",                              :precision => 30, :scale => 30
    t.integer  "placing_user_id"
    t.integer  "position"
    t.integer  "priority"
    t.integer  "shape"
    t.integer  "shipping_id"
    t.integer  "shipping_address_id"
    t.integer  "shipping_method_id"
    t.string   "shipping_number"
    t.string   "url"
    t.string   "user_agent"
    t.string   "referred_by"
    t.string   "utm_campaign"
    t.string   "utm_source"
    t.string   "utm_medium"
    t.string   "utm_term"
    t.string   "utm_content"
    t.integer  "duration"
    t.integer  "length"
    t.integer  "progress"
    t.integer  "resource_id"
    t.string   "resource_name"
    t.string   "resource_type"
    t.integer  "browser_id"
    t.datetime "user_kit_purchase_date"
    t.string   "user_state"
    t.string   "user_region"
    t.datetime "resource_created_at"
    t.datetime "resource_updated_at"
    t.integer  "mobile_popup_id"
    t.string   "client_type"
    t.string   "client_version"
    t.string   "user_slug",               :limit => 4
    t.string   "referring_user_slug",     :limit => 4
    t.integer  "boutique_id"
    t.integer  "product_id"
    t.integer  "referring_user_id"
    t.integer  "taxon_id"
    t.integer  "territory"
    t.integer  "variant_id"
    t.integer  "zip_code_id"
*/

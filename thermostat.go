package main

// Settings and current status for individual thermostat.
type Thermostat struct {
  Identifier string `json:"identifier"`
  Name string `json:"name"`
  ThermostatRev string `json:"thermostatRev"`
  IsRegistered bool `json:"isRegistered"`
  ModelNumber string `json:"modelNumber"`
  Brand string `json:"brand"`
  Features string `json:"features"`
  LastModified string `json:"lastModified"` // TODO: RFC3339 timestamp (UTC?)
  ThermostatTime string `json:"thermostatModified"` // TODO: RFC3339 ts (local?)
  UtcTime string `json:"utcModified"` // TODO: RFC3339 ts (utc?)

  // Thermostat settings
  Settings struct {
    // One of: auto, auxHeatOnly, cool, heat, off
    HvacMode HvacMode `json:"hvacMode"`
    LastServiceDate string `json:"lastServiceDate"` // TODO: rfc3339 date
    ServiceRemindMe bool `json:"serviceRemindMe"`
    MonthsBetweenService int `json:"monthsBetweenService"`
    RemindMeDate string `json:"remindMeDate"` // TODO: rfc3339 date
    Vent string `json:"vent"` // TODO: ?
    VentilatorMinOnTime int `json:"ventilatorMinOnTime"`
    ServiceRemindTechnician bool `json:"serviceremindTechnician"`
    EiLocation string `json:"eiLocation"`
    ColdTempAlert Temperature `json:"coldTempAlert"`
    ColdTempAlertEnabled bool `json:"coldTempAlertEnabled"`
    HotTempAlert Temperature `json:"hotTempAlert"`
    HotTempAlertEnabled bool `json:"hotTempAlertEnabled"`
    CoolStages int `json:"coolStages"`
    HeatStages int `json:"heatStages"`
    MaxSetBack int `json:"maxSetBack"`
    MaxSetForward int `json:"maxSetForward"`
    QuickSaveSetBack int `json:"quickSaveSetBack"`
    QuickSaveSetForward int `json:"quickSaveSetForward"`
    HasHeatPump bool `json:"hasHeatPump"`
    HasForcedAir bool `json:"hasForcedAir"`
    HasBoiler bool `json:"hasBoiler"`
    HasHumidifier bool `json:"hasHumidifier"`
    HasErv bool `json:"hasErv"`
    HasHrv bool `json:"hasHrv"`
    CondensationAvoid bool `json:"condensationAvoid"`
    UseCelcius bool `json:"useCelcius"`
    UseTimeFormat12 bool `json:"useTimeFormat12"`
    Locale string `json:"locale"`
    Humidity string `json:"humidity"`
    HumidifierMode string `json:"humidifierMode"`
    BacklightOnIntensity int `json:"backlightOnIntensity"`
    BacklightSleepIntensity int `json:"backlightSleepIntensity"`
    BacklightOffTime int `json:"backlightOffTime"`
    SoundTickVolume int `json:"soundTickVolume"`
    SoundAlertVolume int `json:"soundAlertVolume"`
    CompressorProtectionMinTime int `json:"compressorProtectionMinTime"`
    CompressorProtectionMinTemp Temperature `json:"compressorProtectionMinTemp"`
    Stage1HeatingDifferentialTemp Temperature `json:"stage1HeatingDifferentialTemp"`
    Stage1CoolingDifferentialTime Temperature `json:"stage1CoolingDifferentialTemp"`
    Stage1HeatingDissipationTime int `json:"stage1HeatingDissipationTime"`
    Stage1CoolingDissipationTime int `json:"stage1CoolingssipationalTemp"`
    HeatPumpReversalOnCool bool `json:"heatPumpReversalOnCool"`
    FanControlRequired bool `json:"fanControlRequired"`
    FanMinOnTime int `json:"fanMinOnTime"`
    HeatCoolMinDelta int `json:"heatCoolMinDelta"`
    TempCorrection int `json:"tempCorrection"`
    HoldAction string `json:"holdAction"`
    HeatPumpGroundWater bool `json:"heatPumpGroundWater"`
    HasElectric bool `json:"hasElectric"`
    HasDehumidifier bool `json:"hasDehumidifier"`
    DehumidiferMode string `json:"dehumidiferMode"`
    DehumidiferLevel int `json:"dehumidiferLevel"`
    DehumidifyWithAC bool `json:"dehumidifyWithAC"`
    DehumidifyOvercoolOffset int `json:"dehumidifyOvercoolOffset"`
    AutoHeatCoolFeatureEnabled bool `json:"autoHeatCoolFeatureEnabled"`
    WifiOfflineAlert bool `json:"wifiOfflineAlert"`
    HeatMinTemp Temperature `json:"heatMinTemp"`
    HeatMaxTemp Temperature `json:"heatMaxTemp"`
    CoolMinTemp Temperature `json:"coolMinTemp"`
    CoolMaxTemp Temperature `json:"coolMaxTemp"`
    HeatRangeHigh Temperature `json:"heatRangeHigh"`
    HeatRangeLow Temperature `json:"heatRangeLow"`
    CoolRangeHigh Temperature `json:"coolRangeHigh"`
    CoolRangeLow Temperature `json:"coolRangeLow"`
    UserAccessCode string `json:"userAccessCode"`
    UserAccessSetting int `json:"userAccessSetting"`
    AuxRuntimeAlert int `json:"auxRuntimeAlert"`
    AuxOutdoorTempAlert Temperature `json:"auxOutdoorTempAlert"`
    AuxMaxOutdoorTemp Temperature `json:"auxMaxOutdoorTemp"`
    AuxRuntimeAlertNotify bool `json:"auxRuntimeAlertNotify"`
    AuxOutdoorTempAlertNotify bool `json:"auxOutdoorTempAlertNotify"`
    AuxRuntimeAlertNotifyTechnician bool `json:"auxRuntimeAlertNotifyTechnician"`
    AuxOutdoorTempAlertNotifyTechnician bool `json:"auxOutdoorTempAlertNotifyTechnician"`
    DisablePreHeating bool `json:"disablePreHeating"`
    DisablePreCooling bool `json:"disablePreCooling"`
    InstallerCodeRequired bool `json:"installerCodeRequired"`
    DrAccept string `json:"drAccept"`
    IsRentalProperty bool `json:"isRentalProperty"`
    UseZoneController bool `json:"useZoneController"`
    RandomStartDelayCool int `json:"randomStartDelayCool"`
    RandomStartDelayHeat int `json:"randomStartDelayHeat"`
    HumidityHighAlert int `json:"humidityHighAlert"`
    HumidityLowAlert int `json:"humidityLowAlert"`
    DisableHeatPumpAlerts bool `json:"disableHeatPumpAlerts"`
    DisableAlertsOnIdt bool `json:"disableAlertsOnIdt"`
    HumidityAlertNotify bool `json:"humidityAlertNotify"`
    HumidityAlertNotifyTechnician bool `json:"humidityAlertNotifyTechnician"`
    TempAlertNotify bool `json:"tempAlertNotify"`
    TempAlertNotifyTechnician bool `json:"tempAlertNotifyTechnician"`
    MonthlyElectricityBillLimit int `json:"monthlyElectricityBillLimit"`
    EnableElectricityBillAlert bool `json:"enableElectricityBillAlert"`
    EnableProjectedElectricityBillAlert bool `json:"enableProjectedElectricityBillAlert"`
    ElectricityBillingDayOfMonth int `json:"electricityBillingDayOfMonth"`
    ElectricityBillCycleMonths int `json:"electricityBillCycleMonths"`
    ElectricityBillStartMonth int `json:"electricityBillStartMonth"`
    VentilatorMinOnTimeHome int `json:"ventilatorMinOnTimeHome"`
    VentilatorMinOnTimeAway int `json:"ventilatorMinOnTimeAway"`
    BacklightOffDuringSleep bool `json:"backlightOffDuringSleep"`
    AutoAway bool `json:"autoAway"`
    SmartCirculation bool `json:"smartCirculation"`
    FollowMeComfort bool `json:"followMeComfort"`
    VentilatorType string `json:"ventilatorType"`
    IsVentilatorTimerOn bool `json:"isVentilatorTimerOn"`
    VentilatorOffDateTime string `json:"ventilatorOffDateTime"` // rfc3339 timestamp (utc?)
    HasUVFilter bool `json:"hasUVFilter"`
    CoolingLockout bool `json:"coolingLockout"`
    VentilatorFreeCooling bool `json:"ventilatorFreeCooling"`
    DehumidifyWhenHeating bool `json:"dehumidifyWhenHeating"`
    VentilatorDehumidify bool `json:"ventilatorDehumidify"`
    GroupRef string `json:"groupRef"`
    GroupName string `json:"groupName"`
    GroupSetting int `json:"groupSetting"`
    FanSpeed string `json:"fanSpeed"`
    DisplayAirQuality bool `json:"displayAirQuality"`
  } `json:"settings"`

  // Current status
  Runtime struct {
    RuntimeRev string `json:"runtimeRev"`
    Connected bool `json:"connected"`
    FirstConnected string `json:"firstConnected"` // rfc3339 timestamp (utc?)
    ConnectDateTime string `json:"connectDateTime"` // rfc3339 timestamp (utc?)
    DisconnectDateTime string `json:"disconnectDateTime"` // rfc3339 timestamp (utc?)
    LastModified string `json:"lastModified"` // rfc3339 timestamp (utc?)
    LastStatusModified string `json:"lastStatusModified"` // rfc3339 timestamp (utc?)
    RuntimeDate string `json:"runtimeDate"` // rfc3339 date
    RuntimeInterval int `json:"runtimeInterval"`
    ActualTemperature Temperature `json:"actualTemperature"`
    ActualHumidity int `json:"actualHumidity"`
    RawTemperature Temperature `json:"rawTemperature"`
    ShowIconMode int `json:"showIconMode"`
    DesiredHeat Temperature `json:"desiredHeat"`
    DesiredCool Temperature `json:"desiredCool"`
    DesiredHumidity int `json:"desiredHumidity"`
    DesiredDehumidity int `json:"desiredDehumidity"`
    DesiredFanMode string `json:"desiredFanMode"`
    ActualVOC int `json:"actualVOC"`
    ActualCO2 int `json:"actualCO2"`
    ActualAQAccuracy int `json:"actualAQAccuracy"`
    ActualAQScore int `json:"actualAQScore"`
    DesiredHeatRange []Temperature `json:"desiredHeatRange"`
    DesiredCoolRange []Temperature `json:"desiredCoolRange"`
  } `json:"runtime"`
}


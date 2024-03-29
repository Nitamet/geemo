// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {lcu} from '../models';

export function ApplyItemSet(arg1:lcu.ItemSet):Promise<void>;

export function ApplyRunes(arg1:lcu.RunePage):Promise<void>;

export function ApplySummonerSpells(arg1:number,arg2:number):Promise<void>;

export function Close():Promise<void>;

export function GetActiveSources():Promise<Array<string>>;

export function GetAssignedRole():Promise<string>;

export function GetAutoImportSetting():Promise<boolean>;

export function GetAutoUpdateSetting():Promise<boolean>;

export function GetCurrentChampion():Promise<number>;

export function GetCurrentVersion():Promise<string>;

export function GetGameMode():Promise<Array<string>>;

export function GetLanguage():Promise<string>;

export function GetShowNativeTitleBarSetting():Promise<boolean>;

export function GetState():Promise<string>;

export function GetSummoner():Promise<lcu.Summoner>;

export function IsUpdateAvailable():Promise<boolean>;

export function Maximize():Promise<void>;

export function Minimize():Promise<void>;

export function OpenLogFolder():Promise<void>;

export function Restart():Promise<void>;

export function SetActiveSources(arg1:Array<string>):Promise<void>;

export function SetAutoImportSetting(arg1:boolean):Promise<void>;

export function SetAutoUpdateSetting(arg1:boolean):Promise<void>;

export function SetLanguage(arg1:string):Promise<void>;

export function SetShowNativeTitleBarSetting(arg1:boolean):Promise<void>;

export function Update():Promise<void>;

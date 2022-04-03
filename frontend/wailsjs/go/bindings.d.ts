import * as models from './models';

export interface go {
  "main": {
    "KoboService": {
		BuildContentIndex(arg1:Array<models.Content>):Promise<models.Content>
		CheckReadwiseConfig():Promise<boolean>
		CheckTokenValidity():Promise<Error>
		CountDeviceBookmarks():Promise<number>
		DetectKobos():Promise<Array<models.Kobo>>
		FindBookOnDevice(arg1:string):Promise<models.Content|Error>
		ForwardToReadwise():Promise<number|Error>
		GetCoverUploadStatus():Promise<boolean>
		GetReadwiseToken():Promise<string>
		GetSelectedKobo():Promise<models.Kobo>
		ListDeviceBookmarks():Promise<Array<models.Bookmark>|Error>
		ListDeviceContent():Promise<Array<models.Content>|Error>
		PromptForLocalDBPath():Promise<Error>
		SelectKobo(arg1:string):Promise<Error>
		SetContext(arg1:models.Context):Promise<void>
		SetCoverUploadStatus(arg1:boolean):Promise<Error>
		SetReadwiseToken(arg1:string):Promise<Error>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}

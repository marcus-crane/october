export interface go {
  "main": {
    "KoboService": {
		BuildContentIndex(arg1:Array<Content>):Promise<Content>
		CheckReadwiseConfig():Promise<boolean>
		CheckTokenValidity():Promise<Error>
		CountDeviceBookmarks():Promise<number>
		DetectKobos():Promise<Array<Kobo>>
		FindBookOnDevice(arg1:string):Promise<Content|Error>
		ForwardToReadwise():Promise<number|Error>
		GetReadwiseToken():Promise<string>
		GetSelectedKobo():Promise<Kobo>
		ListDeviceBookmarks():Promise<Array<Bookmark>|Error>
		ListDeviceContent():Promise<Array<Content>|Error>
		PromptForLocalDBPath():Promise<Error>
		SelectKobo(arg1:string):Promise<Error>
		SetContext(arg1:Context):Promise<void>
		SetReadwiseToken(arg1:string):Promise<Error>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}

export interface go {
  "main": {
    "KoboService": {
		BuildContentIndex(arg1:Array<Content>):Promise<Content>
		CheckReadwiseConfig():Promise<boolean>
		CountDeviceBookmarks():Promise<number>
		DetectKobos():Promise<Array<Kobo>>
		ForwardToReadwise():Promise<number|Error>
		GetSelectedKobo():Promise<Kobo>
		ListDeviceBookmarks():Promise<Array<Bookmark>|Error>
		ListDeviceContent():Promise<Array<Content>|Error>
		PromptForLocalDBPath():Promise<Error>
		SelectKobo(arg1:string):Promise<Error>
		SetContext(arg1:Context):Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}

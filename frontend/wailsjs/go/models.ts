export namespace lcu {
	
	export class Summoner {
	    displayName: string;
	    profileIconId: number;
	
	    static createFrom(source: any = {}) {
	        return new Summoner(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.displayName = source["displayName"];
	        this.profileIconId = source["profileIconId"];
	    }
	}

}


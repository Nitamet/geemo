export namespace lcu {
	
	export class Item {
	    count: number;
	    id: string;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.count = source["count"];
	        this.id = source["id"];
	    }
	}
	export class ItemBlock {
	    items: Item[];
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new ItemBlock(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], Item);
	        this.type = source["type"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ItemSet {
	    associatedChampions: number[];
	    associatedMaps: number[];
	    blocks: ItemBlock[];
	    title: string;
	
	    static createFrom(source: any = {}) {
	        return new ItemSet(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.associatedChampions = source["associatedChampions"];
	        this.associatedMaps = source["associatedMaps"];
	        this.blocks = this.convertValues(source["blocks"], ItemBlock);
	        this.title = source["title"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RunePage {
	    name: string;
	    primaryStyleId: number;
	    subStyleId: number;
	    selectedPerkIds: number[];
	    current: boolean;
	
	    static createFrom(source: any = {}) {
	        return new RunePage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.primaryStyleId = source["primaryStyleId"];
	        this.subStyleId = source["subStyleId"];
	        this.selectedPerkIds = source["selectedPerkIds"];
	        this.current = source["current"];
	    }
	}
	export class Summoner {
	    accountId: number;
	    displayName: string;
	    profileIconId: number;
	    summonerId: number;
	
	    static createFrom(source: any = {}) {
	        return new Summoner(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accountId = source["accountId"];
	        this.displayName = source["displayName"];
	        this.profileIconId = source["profileIconId"];
	        this.summonerId = source["summonerId"];
	    }
	}

}

export namespace lolbuild {
	
	export class Item {
	    id: number;
	    name: string;
	    slug: string;
	    iconUrl: string;
	    isMythic: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.slug = source["slug"];
	        this.iconUrl = source["iconUrl"];
	        this.isMythic = source["isMythic"];
	    }
	}
	export class ItemGroup {
	    items: Item[];
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new ItemGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], Item);
	        this.name = source["name"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SummonerSpell {
	    id: number;
	    iconUrl: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new SummonerSpell(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.iconUrl = source["iconUrl"];
	        this.name = source["name"];
	    }
	}
	export class Rune {
	    id: number;
	    name: string;
	    slug: string;
	    iconUrl: string;
	    path?: Rune;
	
	    static createFrom(source: any = {}) {
	        return new Rune(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.slug = source["slug"];
	        this.iconUrl = source["iconUrl"];
	        this.path = this.convertValues(source["path"], Rune);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Build {
	    name: string;
	    winrate: number;
	    matches: number;
	    primary: Rune;
	    secondary: Rune;
	    selectedPerks: Rune[];
	    summonerSpells: SummonerSpell[];
	    itemGroups: ItemGroup[];
	    mythic: Item;
	
	    static createFrom(source: any = {}) {
	        return new Build(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.winrate = source["winrate"];
	        this.matches = source["matches"];
	        this.primary = this.convertValues(source["primary"], Rune);
	        this.secondary = this.convertValues(source["secondary"], Rune);
	        this.selectedPerks = this.convertValues(source["selectedPerks"], Rune);
	        this.summonerSpells = this.convertValues(source["summonerSpells"], SummonerSpell);
	        this.itemGroups = this.convertValues(source["itemGroups"], ItemGroup);
	        this.mythic = this.convertValues(source["mythic"], Item);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class BuildCollection {
	    builds: Build[];
	    source: string;
	
	    static createFrom(source: any = {}) {
	        return new BuildCollection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.builds = this.convertValues(source["builds"], Build);
	        this.source = source["source"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class BuildSource {
	    slug: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new BuildSource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.slug = source["slug"];
	        this.name = source["name"];
	    }
	}
	export class ChampionName {
	    name: string;
	    slug: string;
	
	    static createFrom(source: any = {}) {
	        return new ChampionName(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.slug = source["slug"];
	    }
	}
	
	
	
	export class RuneTree {
	    name: string;
	    keystones: Rune[];
	    perks: Rune[];
	    iconUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new RuneTree(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.keystones = this.convertValues(source["keystones"], Rune);
	        this.perks = this.convertValues(source["perks"], Rune);
	        this.iconUrl = source["iconUrl"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}


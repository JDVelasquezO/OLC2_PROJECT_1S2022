// Generated from /home/jdvelasquezo/GolandProjects/OLC2_Project1/server/ProjectLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ProjectLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PRINTLN=1, PRINT=2, DECLARAR=3, MUT=4, RSTRING=5, RINTEGER=6, RREAL=7, 
		RBOOLEAN=8, RCHAR=9, RSTR=10, RAS=11, RIF=12, RELSE=13, RMATCH=14, RWHILE=15, 
		RLOOP=16, RFOR=17, RIN=18, RBREAK=19, RCONTINUE=20, RRETURN=21, POWI=22, 
		POWF=23, TOSTRING=24, TOOWNED=25, RMAIN=26, RFN=27, RVECTOR=28, RVEC=29, 
		RVECMayus=30, RNEW=31, REVECTORNEW=32, REVECCAPACITY=33, RSTRUCT=34, RMODULE=35, 
		RPUB=36, MULTILINE=37, INLINE=38, INTEGER=39, FLOAT=40, CHAR=41, STRING=42, 
		BOOLEAN=43, ID=44, EQUAL=45, DOT=46, SEMICOLON=47, COMMA=48, COLON=49, 
		ADMIRATION=50, REFERENCE=51, HERITAGE=52, ARROW=53, RANGE=54, LEFT_PAR=55, 
		RIGHT_PAR=56, LEFT_KEY=57, RIGHT_KEY=58, LEFT_BRACKET=59, RIGHT_BRACKET=60, 
		EQUEAL_EQUAL=61, NOTEQUAL=62, GREATER_THAN=63, LESS_THAN=64, GREATER_EQUALTHAN=65, 
		LESS_EQUEALTHAN=66, MUL=67, DIV=68, MOD=69, ADD=70, SUB=71, AND=72, OR=73, 
		PIPE=74, EQUAL_ARROW=75, UNDERSCORE=76, WHITESPACE=77;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RMATCH", "RWHILE", 
			"RLOOP", "RFOR", "RIN", "RBREAK", "RCONTINUE", "RRETURN", "POWI", "POWF", 
			"TOSTRING", "TOOWNED", "RMAIN", "RFN", "RVECTOR", "RVEC", "RVECMayus", 
			"RNEW", "REVECTORNEW", "REVECCAPACITY", "RSTRUCT", "RMODULE", "RPUB", 
			"MULTILINE", "INLINE", "INTEGER", "FLOAT", "CHAR", "STRING", "BOOLEAN", 
			"ID", "EQUAL", "DOT", "SEMICOLON", "COMMA", "COLON", "ADMIRATION", "REFERENCE", 
			"HERITAGE", "ARROW", "RANGE", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", 
			"LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", 
			"LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL", "DIV", "MOD", 
			"ADD", "SUB", "AND", "OR", "PIPE", "EQUAL_ARROW", "UNDERSCORE", "WHITESPACE", 
			"ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println'", "'print'", "'let'", "'mut'", "'String'", "'i64'", 
			"'f64'", "'bool'", "'char'", "'&str'", "'as'", "'if'", "'else'", "'match'", 
			"'while'", "'loop'", "'for'", "'in'", "'break'", "'continue'", "'return'", 
			"'pow'", "'powf'", "'to_string'", "'to_owned'", "'main'", "'fn'", "'vector'", 
			"'vec'", "'Vec'", "'new'", "'Vec::new()'", "'Vec::with_capacity'", "'struct'", 
			"'mod'", "'pub'", null, null, null, null, null, null, null, null, "'='", 
			"'.'", "';'", "','", "':'", "'!'", "'&'", "'::'", "'->'", "'..'", "'('", 
			"')'", "'{'", "'}'", "'['", "']'", "'=='", "'!='", "'>'", "'<'", "'>='", 
			"'<='", "'*'", "'/'", "'%'", "'+'", "'-'", "'&&'", "'||'", "'|'", "'=>'", 
			"'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RMATCH", "RWHILE", 
			"RLOOP", "RFOR", "RIN", "RBREAK", "RCONTINUE", "RRETURN", "POWI", "POWF", 
			"TOSTRING", "TOOWNED", "RMAIN", "RFN", "RVECTOR", "RVEC", "RVECMayus", 
			"RNEW", "REVECTORNEW", "REVECCAPACITY", "RSTRUCT", "RMODULE", "RPUB", 
			"MULTILINE", "INLINE", "INTEGER", "FLOAT", "CHAR", "STRING", "BOOLEAN", 
			"ID", "EQUAL", "DOT", "SEMICOLON", "COMMA", "COLON", "ADMIRATION", "REFERENCE", 
			"HERITAGE", "ARROW", "RANGE", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", 
			"LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", 
			"LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL", "DIV", "MOD", 
			"ADD", "SUB", "AND", "OR", "PIPE", "EQUAL_ARROW", "UNDERSCORE", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public ProjectLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ProjectLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2O\u020a\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3"+
		"\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3"+
		"\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r\3"+
		"\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3"+
		"\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\23\3\23\3"+
		"\23\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3"+
		"\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\30\3"+
		"\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3"+
		"\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3"+
		"\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3"+
		"\37\3\37\3\37\3\37\3 \3 \3 \3 \3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3\"\3"+
		"\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\""+
		"\3#\3#\3#\3#\3#\3#\3#\3$\3$\3$\3$\3%\3%\3%\3%\3&\3&\3&\3&\6&\u0175\n&"+
		"\r&\16&\u0176\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\6\'\u0182\n\'\r\'\16\'\u0183"+
		"\3\'\3\'\3(\6(\u0189\n(\r(\16(\u018a\3)\6)\u018e\n)\r)\16)\u018f\3)\3"+
		")\6)\u0194\n)\r)\16)\u0195\3*\3*\3*\3*\3+\3+\7+\u019e\n+\f+\16+\u01a1"+
		"\13+\3+\3+\3,\3,\3,\3,\3,\3,\3,\3,\3,\5,\u01ae\n,\3-\3-\7-\u01b2\n-\f"+
		"-\16-\u01b5\13-\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3"+
		"\64\3\64\3\65\3\65\3\65\3\66\3\66\3\66\3\67\3\67\3\67\38\38\39\39\3:\3"+
		":\3;\3;\3<\3<\3=\3=\3>\3>\3>\3?\3?\3?\3@\3@\3A\3A\3B\3B\3B\3C\3C\3C\3"+
		"D\3D\3E\3E\3F\3F\3G\3G\3H\3H\3I\3I\3I\3J\3J\3J\3K\3K\3L\3L\3L\3M\3M\3"+
		"N\6N\u0202\nN\rN\16N\u0203\3N\3N\3O\3O\3O\2\2P\3\3\5\4\7\5\t\6\13\7\r"+
		"\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25"+
		")\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O"+
		")Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081"+
		"B\u0083C\u0085D\u0087E\u0089F\u008bG\u008dH\u008fI\u0091J\u0093K\u0095"+
		"L\u0097M\u0099N\u009bO\u009d\2\3\2\n\3\2\61\61\3\2\f\f\3\2\62;\3\2$$\7"+
		"\2C\\aac|\u00d3\u00d3\u00f3\u00f3\b\2\62;C\\aac|\u00d3\u00d3\u00f3\u00f3"+
		"\5\2\13\f\17\17\"\"\t\2\"#%%--/\60<<BB]_\2\u0211\2\3\3\2\2\2\2\5\3\2\2"+
		"\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21"+
		"\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2"+
		"\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3"+
		"\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3"+
		"\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3"+
		"\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2"+
		"\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2"+
		"Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3"+
		"\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2"+
		"\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2"+
		"\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3"+
		"\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u008d\3\2\2\2\2\u008f\3\2\2\2"+
		"\2\u0091\3\2\2\2\2\u0093\3\2\2\2\2\u0095\3\2\2\2\2\u0097\3\2\2\2\2\u0099"+
		"\3\2\2\2\2\u009b\3\2\2\2\3\u009f\3\2\2\2\5\u00a7\3\2\2\2\7\u00ad\3\2\2"+
		"\2\t\u00b1\3\2\2\2\13\u00b5\3\2\2\2\r\u00bc\3\2\2\2\17\u00c0\3\2\2\2\21"+
		"\u00c4\3\2\2\2\23\u00c9\3\2\2\2\25\u00ce\3\2\2\2\27\u00d3\3\2\2\2\31\u00d6"+
		"\3\2\2\2\33\u00d9\3\2\2\2\35\u00de\3\2\2\2\37\u00e4\3\2\2\2!\u00ea\3\2"+
		"\2\2#\u00ef\3\2\2\2%\u00f3\3\2\2\2\'\u00f6\3\2\2\2)\u00fc\3\2\2\2+\u0105"+
		"\3\2\2\2-\u010c\3\2\2\2/\u0110\3\2\2\2\61\u0115\3\2\2\2\63\u011f\3\2\2"+
		"\2\65\u0128\3\2\2\2\67\u012d\3\2\2\29\u0130\3\2\2\2;\u0137\3\2\2\2=\u013b"+
		"\3\2\2\2?\u013f\3\2\2\2A\u0143\3\2\2\2C\u014e\3\2\2\2E\u0161\3\2\2\2G"+
		"\u0168\3\2\2\2I\u016c\3\2\2\2K\u0170\3\2\2\2M\u017d\3\2\2\2O\u0188\3\2"+
		"\2\2Q\u018d\3\2\2\2S\u0197\3\2\2\2U\u019b\3\2\2\2W\u01ad\3\2\2\2Y\u01af"+
		"\3\2\2\2[\u01b6\3\2\2\2]\u01b8\3\2\2\2_\u01ba\3\2\2\2a\u01bc\3\2\2\2c"+
		"\u01be\3\2\2\2e\u01c0\3\2\2\2g\u01c2\3\2\2\2i\u01c4\3\2\2\2k\u01c7\3\2"+
		"\2\2m\u01ca\3\2\2\2o\u01cd\3\2\2\2q\u01cf\3\2\2\2s\u01d1\3\2\2\2u\u01d3"+
		"\3\2\2\2w\u01d5\3\2\2\2y\u01d7\3\2\2\2{\u01d9\3\2\2\2}\u01dc\3\2\2\2\177"+
		"\u01df\3\2\2\2\u0081\u01e1\3\2\2\2\u0083\u01e3\3\2\2\2\u0085\u01e6\3\2"+
		"\2\2\u0087\u01e9\3\2\2\2\u0089\u01eb\3\2\2\2\u008b\u01ed\3\2\2\2\u008d"+
		"\u01ef\3\2\2\2\u008f\u01f1\3\2\2\2\u0091\u01f3\3\2\2\2\u0093\u01f6\3\2"+
		"\2\2\u0095\u01f9\3\2\2\2\u0097\u01fb\3\2\2\2\u0099\u01fe\3\2\2\2\u009b"+
		"\u0201\3\2\2\2\u009d\u0207\3\2\2\2\u009f\u00a0\7r\2\2\u00a0\u00a1\7t\2"+
		"\2\u00a1\u00a2\7k\2\2\u00a2\u00a3\7p\2\2\u00a3\u00a4\7v\2\2\u00a4\u00a5"+
		"\7n\2\2\u00a5\u00a6\7p\2\2\u00a6\4\3\2\2\2\u00a7\u00a8\7r\2\2\u00a8\u00a9"+
		"\7t\2\2\u00a9\u00aa\7k\2\2\u00aa\u00ab\7p\2\2\u00ab\u00ac\7v\2\2\u00ac"+
		"\6\3\2\2\2\u00ad\u00ae\7n\2\2\u00ae\u00af\7g\2\2\u00af\u00b0\7v\2\2\u00b0"+
		"\b\3\2\2\2\u00b1\u00b2\7o\2\2\u00b2\u00b3\7w\2\2\u00b3\u00b4\7v\2\2\u00b4"+
		"\n\3\2\2\2\u00b5\u00b6\7U\2\2\u00b6\u00b7\7v\2\2\u00b7\u00b8\7t\2\2\u00b8"+
		"\u00b9\7k\2\2\u00b9\u00ba\7p\2\2\u00ba\u00bb\7i\2\2\u00bb\f\3\2\2\2\u00bc"+
		"\u00bd\7k\2\2\u00bd\u00be\78\2\2\u00be\u00bf\7\66\2\2\u00bf\16\3\2\2\2"+
		"\u00c0\u00c1\7h\2\2\u00c1\u00c2\78\2\2\u00c2\u00c3\7\66\2\2\u00c3\20\3"+
		"\2\2\2\u00c4\u00c5\7d\2\2\u00c5\u00c6\7q\2\2\u00c6\u00c7\7q\2\2\u00c7"+
		"\u00c8\7n\2\2\u00c8\22\3\2\2\2\u00c9\u00ca\7e\2\2\u00ca\u00cb\7j\2\2\u00cb"+
		"\u00cc\7c\2\2\u00cc\u00cd\7t\2\2\u00cd\24\3\2\2\2\u00ce\u00cf\7(\2\2\u00cf"+
		"\u00d0\7u\2\2\u00d0\u00d1\7v\2\2\u00d1\u00d2\7t\2\2\u00d2\26\3\2\2\2\u00d3"+
		"\u00d4\7c\2\2\u00d4\u00d5\7u\2\2\u00d5\30\3\2\2\2\u00d6\u00d7\7k\2\2\u00d7"+
		"\u00d8\7h\2\2\u00d8\32\3\2\2\2\u00d9\u00da\7g\2\2\u00da\u00db\7n\2\2\u00db"+
		"\u00dc\7u\2\2\u00dc\u00dd\7g\2\2\u00dd\34\3\2\2\2\u00de\u00df\7o\2\2\u00df"+
		"\u00e0\7c\2\2\u00e0\u00e1\7v\2\2\u00e1\u00e2\7e\2\2\u00e2\u00e3\7j\2\2"+
		"\u00e3\36\3\2\2\2\u00e4\u00e5\7y\2\2\u00e5\u00e6\7j\2\2\u00e6\u00e7\7"+
		"k\2\2\u00e7\u00e8\7n\2\2\u00e8\u00e9\7g\2\2\u00e9 \3\2\2\2\u00ea\u00eb"+
		"\7n\2\2\u00eb\u00ec\7q\2\2\u00ec\u00ed\7q\2\2\u00ed\u00ee\7r\2\2\u00ee"+
		"\"\3\2\2\2\u00ef\u00f0\7h\2\2\u00f0\u00f1\7q\2\2\u00f1\u00f2\7t\2\2\u00f2"+
		"$\3\2\2\2\u00f3\u00f4\7k\2\2\u00f4\u00f5\7p\2\2\u00f5&\3\2\2\2\u00f6\u00f7"+
		"\7d\2\2\u00f7\u00f8\7t\2\2\u00f8\u00f9\7g\2\2\u00f9\u00fa\7c\2\2\u00fa"+
		"\u00fb\7m\2\2\u00fb(\3\2\2\2\u00fc\u00fd\7e\2\2\u00fd\u00fe\7q\2\2\u00fe"+
		"\u00ff\7p\2\2\u00ff\u0100\7v\2\2\u0100\u0101\7k\2\2\u0101\u0102\7p\2\2"+
		"\u0102\u0103\7w\2\2\u0103\u0104\7g\2\2\u0104*\3\2\2\2\u0105\u0106\7t\2"+
		"\2\u0106\u0107\7g\2\2\u0107\u0108\7v\2\2\u0108\u0109\7w\2\2\u0109\u010a"+
		"\7t\2\2\u010a\u010b\7p\2\2\u010b,\3\2\2\2\u010c\u010d\7r\2\2\u010d\u010e"+
		"\7q\2\2\u010e\u010f\7y\2\2\u010f.\3\2\2\2\u0110\u0111\7r\2\2\u0111\u0112"+
		"\7q\2\2\u0112\u0113\7y\2\2\u0113\u0114\7h\2\2\u0114\60\3\2\2\2\u0115\u0116"+
		"\7v\2\2\u0116\u0117\7q\2\2\u0117\u0118\7a\2\2\u0118\u0119\7u\2\2\u0119"+
		"\u011a\7v\2\2\u011a\u011b\7t\2\2\u011b\u011c\7k\2\2\u011c\u011d\7p\2\2"+
		"\u011d\u011e\7i\2\2\u011e\62\3\2\2\2\u011f\u0120\7v\2\2\u0120\u0121\7"+
		"q\2\2\u0121\u0122\7a\2\2\u0122\u0123\7q\2\2\u0123\u0124\7y\2\2\u0124\u0125"+
		"\7p\2\2\u0125\u0126\7g\2\2\u0126\u0127\7f\2\2\u0127\64\3\2\2\2\u0128\u0129"+
		"\7o\2\2\u0129\u012a\7c\2\2\u012a\u012b\7k\2\2\u012b\u012c\7p\2\2\u012c"+
		"\66\3\2\2\2\u012d\u012e\7h\2\2\u012e\u012f\7p\2\2\u012f8\3\2\2\2\u0130"+
		"\u0131\7x\2\2\u0131\u0132\7g\2\2\u0132\u0133\7e\2\2\u0133\u0134\7v\2\2"+
		"\u0134\u0135\7q\2\2\u0135\u0136\7t\2\2\u0136:\3\2\2\2\u0137\u0138\7x\2"+
		"\2\u0138\u0139\7g\2\2\u0139\u013a\7e\2\2\u013a<\3\2\2\2\u013b\u013c\7"+
		"X\2\2\u013c\u013d\7g\2\2\u013d\u013e\7e\2\2\u013e>\3\2\2\2\u013f\u0140"+
		"\7p\2\2\u0140\u0141\7g\2\2\u0141\u0142\7y\2\2\u0142@\3\2\2\2\u0143\u0144"+
		"\7X\2\2\u0144\u0145\7g\2\2\u0145\u0146\7e\2\2\u0146\u0147\7<\2\2\u0147"+
		"\u0148\7<\2\2\u0148\u0149\7p\2\2\u0149\u014a\7g\2\2\u014a\u014b\7y\2\2"+
		"\u014b\u014c\7*\2\2\u014c\u014d\7+\2\2\u014dB\3\2\2\2\u014e\u014f\7X\2"+
		"\2\u014f\u0150\7g\2\2\u0150\u0151\7e\2\2\u0151\u0152\7<\2\2\u0152\u0153"+
		"\7<\2\2\u0153\u0154\7y\2\2\u0154\u0155\7k\2\2\u0155\u0156\7v\2\2\u0156"+
		"\u0157\7j\2\2\u0157\u0158\7a\2\2\u0158\u0159\7e\2\2\u0159\u015a\7c\2\2"+
		"\u015a\u015b\7r\2\2\u015b\u015c\7c\2\2\u015c\u015d\7e\2\2\u015d\u015e"+
		"\7k\2\2\u015e\u015f\7v\2\2\u015f\u0160\7{\2\2\u0160D\3\2\2\2\u0161\u0162"+
		"\7u\2\2\u0162\u0163\7v\2\2\u0163\u0164\7t\2\2\u0164\u0165\7w\2\2\u0165"+
		"\u0166\7e\2\2\u0166\u0167\7v\2\2\u0167F\3\2\2\2\u0168\u0169\7o\2\2\u0169"+
		"\u016a\7q\2\2\u016a\u016b\7f\2\2\u016bH\3\2\2\2\u016c\u016d\7r\2\2\u016d"+
		"\u016e\7w\2\2\u016e\u016f\7d\2\2\u016fJ\3\2\2\2\u0170\u0171\7\61\2\2\u0171"+
		"\u0172\7,\2\2\u0172\u0174\3\2\2\2\u0173\u0175\n\2\2\2\u0174\u0173\3\2"+
		"\2\2\u0175\u0176\3\2\2\2\u0176\u0174\3\2\2\2\u0176\u0177\3\2\2\2\u0177"+
		"\u0178\3\2\2\2\u0178\u0179\7,\2\2\u0179\u017a\7\61\2\2\u017a\u017b\3\2"+
		"\2\2\u017b\u017c\b&\2\2\u017cL\3\2\2\2\u017d\u017e\7\61\2\2\u017e\u017f"+
		"\7\61\2\2\u017f\u0181\3\2\2\2\u0180\u0182\n\3\2\2\u0181\u0180\3\2\2\2"+
		"\u0182\u0183\3\2\2\2\u0183\u0181\3\2\2\2\u0183\u0184\3\2\2\2\u0184\u0185"+
		"\3\2\2\2\u0185\u0186\b\'\2\2\u0186N\3\2\2\2\u0187\u0189\t\4\2\2\u0188"+
		"\u0187\3\2\2\2\u0189\u018a\3\2\2\2\u018a\u0188\3\2\2\2\u018a\u018b\3\2"+
		"\2\2\u018bP\3\2\2\2\u018c\u018e\t\4\2\2\u018d\u018c\3\2\2\2\u018e\u018f"+
		"\3\2\2\2\u018f\u018d\3\2\2\2\u018f\u0190\3\2\2\2\u0190\u0191\3\2\2\2\u0191"+
		"\u0193\7\60\2\2\u0192\u0194\t\4\2\2\u0193\u0192\3\2\2\2\u0194\u0195\3"+
		"\2\2\2\u0195\u0193\3\2\2\2\u0195\u0196\3\2\2\2\u0196R\3\2\2\2\u0197\u0198"+
		"\7)\2\2\u0198\u0199\n\5\2\2\u0199\u019a\7)\2\2\u019aT\3\2\2\2\u019b\u019f"+
		"\7$\2\2\u019c\u019e\n\5\2\2\u019d\u019c\3\2\2\2\u019e\u01a1\3\2\2\2\u019f"+
		"\u019d\3\2\2\2\u019f\u01a0\3\2\2\2\u01a0\u01a2\3\2\2\2\u01a1\u019f\3\2"+
		"\2\2\u01a2\u01a3\7$\2\2\u01a3V\3\2\2\2\u01a4\u01a5\7v\2\2\u01a5\u01a6"+
		"\7t\2\2\u01a6\u01a7\7w\2\2\u01a7\u01ae\7g\2\2\u01a8\u01a9\7h\2\2\u01a9"+
		"\u01aa\7c\2\2\u01aa\u01ab\7n\2\2\u01ab\u01ac\7u\2\2\u01ac\u01ae\7g\2\2"+
		"\u01ad\u01a4\3\2\2\2\u01ad\u01a8\3\2\2\2\u01aeX\3\2\2\2\u01af\u01b3\t"+
		"\6\2\2\u01b0\u01b2\t\7\2\2\u01b1\u01b0\3\2\2\2\u01b2\u01b5\3\2\2\2\u01b3"+
		"\u01b1\3\2\2\2\u01b3\u01b4\3\2\2\2\u01b4Z\3\2\2\2\u01b5\u01b3\3\2\2\2"+
		"\u01b6\u01b7\7?\2\2\u01b7\\\3\2\2\2\u01b8\u01b9\7\60\2\2\u01b9^\3\2\2"+
		"\2\u01ba\u01bb\7=\2\2\u01bb`\3\2\2\2\u01bc\u01bd\7.\2\2\u01bdb\3\2\2\2"+
		"\u01be\u01bf\7<\2\2\u01bfd\3\2\2\2\u01c0\u01c1\7#\2\2\u01c1f\3\2\2\2\u01c2"+
		"\u01c3\7(\2\2\u01c3h\3\2\2\2\u01c4\u01c5\7<\2\2\u01c5\u01c6\7<\2\2\u01c6"+
		"j\3\2\2\2\u01c7\u01c8\7/\2\2\u01c8\u01c9\7@\2\2\u01c9l\3\2\2\2\u01ca\u01cb"+
		"\7\60\2\2\u01cb\u01cc\7\60\2\2\u01ccn\3\2\2\2\u01cd\u01ce\7*\2\2\u01ce"+
		"p\3\2\2\2\u01cf\u01d0\7+\2\2\u01d0r\3\2\2\2\u01d1\u01d2\7}\2\2\u01d2t"+
		"\3\2\2\2\u01d3\u01d4\7\177\2\2\u01d4v\3\2\2\2\u01d5\u01d6\7]\2\2\u01d6"+
		"x\3\2\2\2\u01d7\u01d8\7_\2\2\u01d8z\3\2\2\2\u01d9\u01da\7?\2\2\u01da\u01db"+
		"\7?\2\2\u01db|\3\2\2\2\u01dc\u01dd\7#\2\2\u01dd\u01de\7?\2\2\u01de~\3"+
		"\2\2\2\u01df\u01e0\7@\2\2\u01e0\u0080\3\2\2\2\u01e1\u01e2\7>\2\2\u01e2"+
		"\u0082\3\2\2\2\u01e3\u01e4\7@\2\2\u01e4\u01e5\7?\2\2\u01e5\u0084\3\2\2"+
		"\2\u01e6\u01e7\7>\2\2\u01e7\u01e8\7?\2\2\u01e8\u0086\3\2\2\2\u01e9\u01ea"+
		"\7,\2\2\u01ea\u0088\3\2\2\2\u01eb\u01ec\7\61\2\2\u01ec\u008a\3\2\2\2\u01ed"+
		"\u01ee\7\'\2\2\u01ee\u008c\3\2\2\2\u01ef\u01f0\7-\2\2\u01f0\u008e\3\2"+
		"\2\2\u01f1\u01f2\7/\2\2\u01f2\u0090\3\2\2\2\u01f3\u01f4\7(\2\2\u01f4\u01f5"+
		"\7(\2\2\u01f5\u0092\3\2\2\2\u01f6\u01f7\7~\2\2\u01f7\u01f8\7~\2\2\u01f8"+
		"\u0094\3\2\2\2\u01f9\u01fa\7~\2\2\u01fa\u0096\3\2\2\2\u01fb\u01fc\7?\2"+
		"\2\u01fc\u01fd\7@\2\2\u01fd\u0098\3\2\2\2\u01fe\u01ff\7a\2\2\u01ff\u009a"+
		"\3\2\2\2\u0200\u0202\t\b\2\2\u0201\u0200\3\2\2\2\u0202\u0203\3\2\2\2\u0203"+
		"\u0201\3\2\2\2\u0203\u0204\3\2\2\2\u0204\u0205\3\2\2\2\u0205\u0206\bN"+
		"\2\2\u0206\u009c\3\2\2\2\u0207\u0208\7^\2\2\u0208\u0209\t\t\2\2\u0209"+
		"\u009e\3\2\2\2\f\2\u0176\u0183\u018a\u018f\u0195\u019f\u01ad\u01b3\u0203"+
		"\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}